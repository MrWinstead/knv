package cluster

import "context"

type MembershipManager interface {
	Run(ctx context.Context) error
	MembershipReported() <-chan bool
}
