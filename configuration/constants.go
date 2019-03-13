package configuration

const (
	// KeyDirectionNameRootDatabase is the root where the database tables and
	// indexes will be stored
	KeyDirectoryNameRootDatabase = "DIR_NAME_ROOT_DATABASE"

	// KeyPortGRPC is the TCP port upon which knv should expose the gRPC service
	KeyPortGRPC = "PORT_GRPC"

	KeyObservationBacklogMax = "OBSERVATION_BACKLOG_MAX"
)
