package cluster

import (
	"context"
	"errors"
	"github.com/blang/semver"
)

var (
	ErrorNotLeader            = errors.New("not a leader for operation")
	ErrorClusterAlreadyExists = errors.New("cluster already exists")
	ErrorClusterTooNew        = errors.New("cluster version listed as too new")
	ErrorClusterTooOld        = errors.New("cluster version listed as too old")
)

type MembershipManager interface {
	// Run ensures that the node will maintain their membership in the cluster
	Run(ctx context.Context) error

	// MembershipReported provides a method by which to understand whether the
	// node successfully reported as a cluster member
	MembershipReported() <-chan bool
}

// Topology holds metadata regarding the cluster and cluster topology
type Topology struct {
	// Name of the cluster to delineate from other clusters which may share the
	// same infrastructure
	Name string

	// Version of the cluster and associated protocols
	Version semver.Version

	// ShardCount is the number of shards into which the cluster will divide the
	// keyspace
	ShardCount uint
}

type TopographyManager interface {
	// Run will vie for election to initialize the cluster
	Run(ctx context.Context) error

	// IsLeader will report whether the node is a leader to initialize the
	// cluster
	IsLeader() <-chan bool

	// SubmitClusterInformation will ensure that the cluster information is
	// recorded
	//
	// returns ErrorNotLeader if the node is not a leader
	SubmitClusterInformation(ctx context.Context, info *Topology) error
}
