package cluster

import (
	"bytes"
	"context"
	"encoding/json"
	"sync"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/pkg/errors"

	"github.com/mrwinstead/knv/observation"
)

const (
	sinkNameEtcdInitializerObservation = "EtcdTopographyObservation"
)

var (
	_ TopographyManager = &EtcdTopographyManager{}
)

type EtcdTopographyObservation struct {
	SavedClusterTopology Topology
	Error                error
}

type EtcdTopographyElectedValue struct {
	Identity ReportedIdentity
}

type EtcdTopographyManager struct {
	client   *clientv3.Client
	session  *concurrency.Session
	identity ReportedIdentity

	obsSink     observation.Sink
	pathAdvisor PathAdvisor

	isLeader        bool
	isleaderChan    chan bool
	leadershipAwait sync.WaitGroup
}

func NewEtcdInitializer(client *clientv3.Client, session *concurrency.Session,
	obsMgr observation.Manager, pathAdvisor PathAdvisor,
	identity ReportedIdentity) TopographyManager {
	created := &EtcdTopographyManager{
		client:      client,
		identity:    identity,
		session:     session,
		pathAdvisor: pathAdvisor,

		obsSink:      obsMgr.Sink(sinkNameEtcdInitializerObservation),
		isleaderChan: make(chan bool),
	}

	return created
}

func (e *EtcdTopographyManager) Run(ctx context.Context) error {
	if nil == e.isleaderChan {
		e.isleaderChan = make(chan bool)
	}

	// report leadership
	go func() {
		isLeader := <-e.isleaderChan
		for {
			select {
			case e.isleaderChan <- isLeader:
				continue
			case <-ctx.Done():
				return
			}
		}
	}()

	electionPath := e.pathAdvisor.ExpandPath(PathPrefixInitializerElection)
	election := concurrency.NewElection(e.session, electionPath)

	electedValue := EtcdTopographyElectedValue{Identity: e.identity}
	serializedElectedValue := bytes.Buffer{}
	// error should not happen
	json.NewEncoder(&serializedElectedValue).Encode(electedValue)

	electionErr := election.Campaign(ctx, serializedElectedValue.String())
	if nil != electionErr && context.Canceled != electionErr {
		err := errors.Wrap(electionErr, "could not campaign for"+
			"election to initialize cluster")
		return err
	} else if context.Canceled == electionErr {
		return context.Canceled
	}

	e.isleaderChan <- true

	<-ctx.Done()

	return nil
}

func (e *EtcdTopographyManager) IsLeader() <-chan bool {
	return e.isleaderChan
}

func (e *EtcdTopographyManager) SubmitClusterInformation(ctx context.Context,
	info *Topology) error {
	if nil != e.isleaderChan {
		if isLeader := <-e.isleaderChan; !isLeader {
			return ErrorNotLeader
		}
	}

	informationPath := e.pathAdvisor.ExpandPath(PathClusterInformation)

	preexistingClusterInformation, getErr := e.client.Get(ctx, informationPath)
	if nil != getErr && context.Canceled != getErr {
		err := errors.Wrap(getErr, "could not determine if cluster "+
			"already exists")
		return err
	} else if context.Canceled == getErr {
		return context.Canceled
	}

	if 0 < preexistingClusterInformation.Count {
		return ErrorClusterAlreadyExists
	}

	serializedTopography := bytes.Buffer{}
	// error should not happen
	json.NewEncoder(&serializedTopography).Encode(info)

	_, putErr := e.client.Put(ctx, informationPath,
		serializedTopography.String())
	if nil != putErr {
		err := errors.Wrap(putErr, "unable to emplace cluster "+
			"information")
		return err
	}

	return nil
}
