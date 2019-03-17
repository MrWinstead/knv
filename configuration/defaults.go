package configuration

import (
	"time"

	"github.com/spf13/viper"
)

const (
	// DefaultPortGRPC is the default port for for the gRPC service
	DefaultPortGRPC = 8080

	// DefaultObservationBacklogMax is the maximum number of observations to
	// buffer before blocking observation publications
	DefaultObservationBacklogMax = 100

	// DefaultLivenessLeaseTime of 5 seconds
	DefaultLivenessLeaseTime = 5 * time.Second

	// DefaultClusterInformationPathPrefix is a fully-qualified path for this
	// project
	DefaultClusterInformationPathPrefix = "/github.com/mrwinstead/knv/"

	// DefaultBackingServiceEtcd defaults to localhost on default port
	DefaultBackingServiceEtcd = "127.0.0.1:2379"

	// DefaultClusterName something descriptive
	DefaultClusterName = "default"

	// DefaultClusterShardCount defaults to 8 shards
	DefaultClusterShardCount = 8
)

var (
	defaultsMap = map[string]interface{}{
		KeyPortGRPC:                     DefaultPortGRPC,
		KeyObservationBacklogMax:        DefaultObservationBacklogMax,
		KeyMembershipLeaseTime:          DefaultLivenessLeaseTime,
		KeyClusterInformationPathPrefix: DefaultClusterInformationPathPrefix,
		KeyBackingServiceEtcd:           DefaultBackingServiceEtcd,
		KeyClusterName:                  DefaultClusterName,
		KeyClusterShardCount:            DefaultClusterShardCount,
	}
)

func PopulateDefaultValues(v *viper.Viper) {
	for key, value := range defaultsMap {
		v.SetDefault(key, value)
	}
}
