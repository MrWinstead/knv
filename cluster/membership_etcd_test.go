package cluster

import (
	"context"
	"github.com/mrwinstead/knv/configuration"
	"github.com/spf13/viper"
	"io"
	"math/rand"
	"path"
	"testing"
	"time"

	"github.com/etcd-io/etcd/clientv3"
	"github.com/oklog/ulid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"

	"github.com/mrwinstead/knv/observation"
)

const (
	defaultTestMaxRuntime = 20 * time.Second
)

var (
	_ suite.TestingSuite = &etcdLivenessReporterTestSuite{}
)

type etcdLivenessReporterTestSuite struct {
	suite.Suite

	reporter *EtcdMembershipReporter

	obsMgr observation.Manager

	client   *clientv3.Client
	ctx      context.Context
	idSource io.Reader
}

func (e *etcdLivenessReporterTestSuite) SetupSuite() {
	e.obsMgr = observation.NewChannelManager(uint(100))

	endpoints := viper.GetStringSlice(configuration.KeyBackingServiceEtcd)
	client, buildErr := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	})
	if nil != buildErr {
		err := errors.Wrap(buildErr, "could not create etcdv3 client")
		e.Fail(err.Error())
	}
	e.client = client
	e.idSource = ulid.Monotonic(rand.New(rand.NewSource(rand.Int63())), 0)
	pathPrefix := path.Join("test",
		ulid.MustNew(ulid.Now(), e.idSource).String())
	identity := ulid.MustNew(ulid.Now(), e.idSource)

	pa := NewEtcdPathAdvisor(pathPrefix)

	reporter, reporterBuildErr := NewEtcdMembershipReporter(e.client, e.obsMgr,
		ReportedIdentity(identity.String()), pa, 1*time.Second)
	e.NoError(reporterBuildErr)
	e.reporter = reporter.(*EtcdMembershipReporter)
}

func (e *etcdLivenessReporterTestSuite) SetupTest() {
	e.ctx, _ = context.WithTimeout(context.Background(), defaultTestMaxRuntime)
}

func (e *etcdLivenessReporterTestSuite) TestSetLeaseAlgorithm() {
	clientID := ulid.MustNew(ulid.Now(), e.idSource)
	value := ulid.MustNew(ulid.Now(), e.idSource)

	lease, leaseErr := e.client.Grant(context.Background(), 30)
	e.NoError(leaseErr)

	leasePath := "/test/" + clientID.String()

	putResponse, putErr := e.client.Put(e.ctx, leasePath, value.String(),
		clientv3.WithLease(lease.ID))
	e.NoError(putErr)
	e.NotEmpty(putResponse.Header.String())

	getResponse, getErr := e.client.Get(e.ctx, leasePath)
	e.NoError(getErr, "cold not get path", leasePath)
	e.Equal(int64(1), getResponse.Count)

	listResponse, listErr := e.client.Get(e.ctx, "/", clientv3.WithPrefix())
	e.NoError(listErr)
	e.NotEqual(0, listResponse.Count, "no leases found")
}

func (e *etcdLivenessReporterTestSuite) TestLeaseRenew() {
	reporterCtx, cancel := context.WithCancel(e.ctx)
	defer cancel()

	go func() {
		runErr := e.reporter.Run(reporterCtx)
		if context.Canceled == runErr {
			return
		}
		e.NoError(runErr)
	}()

	isMember := <-e.reporter.MembershipReported()
	e.True(isMember, "membership was not reported")
}

func TestEtcdLivenessReporter(t *testing.T) {
	viper.AutomaticEnv()
	endpoints := viper.GetStringSlice(configuration.KeyBackingServiceEtcd)
	if nil != endpoints && 0 == len(endpoints) {
		t.Skip(configuration.KeyBackingServiceEtcd + " environment variable not set")
	}
	suite.Run(t, new(etcdLivenessReporterTestSuite))
}
