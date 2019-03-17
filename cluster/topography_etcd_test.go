package cluster

import (
	"context"
	"io"
	"math/rand"
	"path"
	"testing"
	"time"

	"github.com/blang/semver"
	"github.com/brianvoe/gofakeit"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/oklog/ulid"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"

	"github.com/mrwinstead/knv/configuration"
	"github.com/mrwinstead/knv/observation"
)

type etcdInitializerTestSuite struct {
	suite.Suite

	obsMgr observation.Manager

	client    *clientv3.Client
	session   *concurrency.Session
	ctx       context.Context
	ctxCancel context.CancelFunc
	idSource  io.Reader

	etcdSuitePathPrefix string
	etcdTestPathPrefix  string

	identity    string
	initializer *EtcdTopographyManager
}

func (e *etcdInitializerTestSuite) SetupSuite() {
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

	session, sessionBuildErr := concurrency.NewSession(client)
	e.NoError(sessionBuildErr)
	e.session = session

	e.idSource = ulid.Monotonic(rand.New(rand.NewSource(rand.Int63())), 0)

	e.etcdSuitePathPrefix = path.Join("/test", gofakeit.Name())
}

func (e *etcdInitializerTestSuite) TearDownSuite() {
	_, deleteErr := e.client.Delete(context.Background(), e.etcdSuitePathPrefix,
		clientv3.WithPrefix())
	e.NoError(deleteErr)
}

func (e *etcdInitializerTestSuite) SetupTest() {
	e.ctx, e.ctxCancel = context.WithTimeout(context.Background(), defaultTestMaxRuntime)
	identity := ulid.MustNew(ulid.Now(), e.idSource)
	e.identity = identity.String()
	e.etcdTestPathPrefix = path.Join(e.etcdSuitePathPrefix, e.identity)
	pathAdvisor := NewEtcdPathAdvisor(e.etcdSuitePathPrefix)
	e.initializer = NewEtcdInitializer(e.client, e.session, e.obsMgr, pathAdvisor,
		ReportedIdentity(e.identity)).(*EtcdTopographyManager)

	go func() {
		runErr := e.initializer.Run(e.ctx)
		if context.Canceled == runErr {
			return
		}
		e.NoError(runErr)
	}()
}

func (e *etcdInitializerTestSuite) TearDownTest() {
	e.ctxCancel()
}

func (e *etcdInitializerTestSuite) TestIsLeader_NotStarted() {
	// Drain the automatically-started initializer
	e.ctxCancel()
	select {
	case <-e.initializer.IsLeader():
	default:
	}

	var isLeaderValue *bool
	select {
	case isLeader := <-e.initializer.IsLeader():
		isLeaderValue = &isLeader
	default:
	}
	e.Nil(isLeaderValue, "did not expect to have leadership "+
		"value without campaigning")
}

func (e *etcdInitializerTestSuite) TestIsLeader_Started() {
	var isLeaderValue bool
	select {
	case isLeader := <-e.initializer.IsLeader():
		isLeaderValue = isLeader
	case <-e.ctx.Done():
		e.Fail("did not complete leadership in time")
	}
	e.True(isLeaderValue, "expected be leader")
}

func (e *etcdInitializerTestSuite) TestInitLeadership() {
	initializerCtx, cancel := context.WithCancel(e.ctx)
	defer cancel()

	go func() {
		runErr := e.initializer.Run(initializerCtx)
		if context.Canceled == runErr {
			return
		}
		e.NoError(runErr)
	}()

	select {
	case isLeader := <-e.initializer.isleaderChan:
		e.True(isLeader, "was not elected leader")
	case <-e.ctx.Done():
		e.Fail("was never elected leader before context cancelled")
	}

	topo := Topology{
		Name:       gofakeit.Name(),
		Version:    semver.Version{Major: rand.Uint64(), Minor: rand.Uint64()},
		ShardCount: configuration.DefaultClusterShardCount,
	}
	submissionErr := e.initializer.SubmitClusterInformation(e.ctx, &topo)
	e.NoError(submissionErr, "error encountered while "+
		"submitting topography")

	serializedTopography, getErr := e.client.Get(e.ctx,
		e.initializer.pathAdvisor.ExpandPath(PathClusterInformation))
	e.NoError(getErr, "could not fetch cluster information")
	e.NotEmpty(serializedTopography.Kvs)
	e.Contains(string(serializedTopography.Kvs[0].Value), topo.Name)
}

func TestEtcdTopographyManager(t *testing.T) {
	viper.AutomaticEnv()
	endpoints := viper.GetStringSlice(configuration.KeyBackingServiceEtcd)
	if nil != endpoints && 0 == len(endpoints) {
		t.Skip(configuration.KeyBackingServiceEtcd + " environment variable not set")
	}
	suite.Run(t, new(etcdInitializerTestSuite))
}
