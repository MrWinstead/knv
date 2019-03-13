package configuration

import "github.com/spf13/viper"

const (
	// DefaultPortGRPC is the default port for for the gRPC service
	DefaultPortGRPC = 8080

	// DefaultObservationBacklogMax is the maximum number of observations to
	// buffer before blocking observation publications
	DefaultObservationBacklogMax = 100
)

var (
	defaultsMap = map[string]interface{}{
		KeyPortGRPC:              DefaultPortGRPC,
		KeyObservationBacklogMax: DefaultObservationBacklogMax,
	}
)

func PopulateDefaultValues(v *viper.Viper) {
	for key, value := range defaultsMap {
		v.SetDefault(key, value)
	}
}
