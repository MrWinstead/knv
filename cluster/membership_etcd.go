package cluster

import (
	"context"
	"runtime"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"

	"github.com/mrwinstead/knv/observation"
)

const (
	sinkNameEtcdMembershipReporter = "EtcdMembershipReporter"
)

var (
	_ MembershipManager = &EtcdMembershipReporter{}
)

type ReportedIdentity string

type EtcdMembershipObservable struct {
	TTL       time.Duration
	ClusterID uint64
	Path      string

	Error    error
	JoinTime time.Time

	StopBecauseContextEnded bool
}

type EtcdMembershipReporter struct {
	client                 *clientv3.Client
	currentMembershipLease clientv3.LeaseID
	membershipReportedChan chan bool

	pathAdvisor PathAdvisor

	leaseTime time.Duration

	reporterCancel context.CancelFunc
	identity       ReportedIdentity

	sinkMembership observation.Sink
}

func NewEtcdMembershipReporter(client *clientv3.Client, mgr observation.Manager,
	identity ReportedIdentity, advisor PathAdvisor,
	membershipLeaseTime time.Duration) (MembershipManager, error) {
	created := &EtcdMembershipReporter{
		pathAdvisor:            advisor,
		client:                 client,
		identity:               identity,
		sinkMembership:         mgr.Sink(sinkNameEtcdMembershipReporter),
		membershipReportedChan: make(chan bool),

		leaseTime: membershipLeaseTime,
	}

	runtime.SetFinalizer(created, func(_ interface{}) {
		close(created.membershipReportedChan)
	})

	if 0 == created.leaseTime {
		err := errors.New("lease time must be specified")
		return nil, err
	}

	return created, nil
}

func (lr *EtcdMembershipReporter) Run(ctx context.Context) error {

	// Start channel cleanup functionality
	go func() {
		select {
		case <-ctx.Done():
			return
		}
	}()

	joinTime := time.Now()
	functionObservable := &EtcdMembershipObservable{
		JoinTime: joinTime,
	}
	defer lr.sinkMembership.Submit(nil, functionObservable)

	keepAliveChan, keepAliveErr := lr.InitializeLease(ctx)
	if nil != keepAliveErr {
		functionObservable.Error = keepAliveErr
		return keepAliveErr
	}

	membershipEntryErr := lr.EnsureMembershipEntry(ctx)
	if nil != membershipEntryErr {
		return membershipEntryErr
	}

	// report we're a member for now onwards
	membershipReportedCtx, membershipReportedCtxCancel := context.WithCancel(ctx)
	defer membershipReportedCtxCancel()
	go func() {
		for {
			select {
			case lr.membershipReportedChan <- true:
				continue
			case <-membershipReportedCtx.Done():
				return
			}
		}
	}()

	for {
		membershipShouldBeRestored := false
		obs := &EtcdMembershipObservable{}
		*obs = *functionObservable

		select {
		case keepAliveMessage := <-keepAliveChan:
			if nil == keepAliveMessage { // then the etcd server has done away
				err := errors.New("etcd server has gone away")
				obs.Error = err
				membershipShouldBeRestored = true
			}
			obs.ClusterID = keepAliveMessage.ClusterId
			obs.TTL = time.Duration(keepAliveMessage.TTL) * time.Second
		case <-ctx.Done():
			obs.StopBecauseContextEnded = true
			return context.Canceled
		}

		lr.sinkMembership.Submit(nil, obs)

		if membershipShouldBeRestored {
			// handle recovery case
			panic("restoring cluster membership not implemented")
		}
	}

	return nil // never reached
}

func (lr *EtcdMembershipReporter) EnsureMembershipEntry(ctx context.Context) error {
	obs := &EtcdMembershipObservable{}
	defer lr.sinkMembership.Submit(nil, obs)

	membershipPath := lr.pathAdvisor.ExpandPath(PathMembers,
		string(lr.identity))
	obs.Path = membershipPath

	_, putErr := lr.client.Put(ctx, membershipPath, "",
		clientv3.WithLease(lr.currentMembershipLease))
	if nil != putErr {
		err := errors.Wrap(putErr, "could not emplace membership "+
			"information")
		return err
	}

	return nil
}

func (lr *EtcdMembershipReporter) MembershipReported() <-chan bool {
	return lr.membershipReportedChan
}

func (lr *EtcdMembershipReporter) InitializeLease(ctx context.Context) (
	<-chan *clientv3.LeaseKeepAliveResponse, error) {
	leaseID, leaseErr := lr.RegisterNewMembershipLease(ctx)
	if nil != leaseErr {
		return nil, leaseErr
	}
	lr.currentMembershipLease = leaseID

	keepAliveChan, keepAliveErr := lr.client.KeepAlive(ctx, leaseID)
	if nil != keepAliveErr {
		err := errors.Wrap(keepAliveErr, "could not create "+
			"continuous cluster membership lease manager")
		return nil, err
	}

	return keepAliveChan, nil
}

func (lr *EtcdMembershipReporter) RegisterNewMembershipLease(ctx context.Context,
) (clientv3.LeaseID, error) {
	lease, leaseErr := lr.client.Grant(ctx, int64(lr.leaseTime/time.Second))
	if nil != leaseErr {
		err := errors.Wrap(leaseErr, "could not establish "+
			"cluster membership lease")
		return 0, err
	}

	return lease.ID, nil
}
