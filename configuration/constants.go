package configuration

const (
	// KeyDirectionNameRootDatabase is the root where the database tables and
	// indexes will be stored
	KeyDirectoryNameRootDatabase = "DIR_NAME_ROOT_DATABASE"

	// KeyPortGRPC is the TCP port upon which knv should expose the gRPC service
	KeyPortGRPC = "PORT_GRPC"

	// KeyObservationBacklogMax is the number of observations to backlog before
	// hanging reporters
	KeyObservationBacklogMax = "OBSERVATION_BACKLOG_MAX"

	// KeyMembershipLeaseTime is the amount of time which a node will be
	// listed as alive
	KeyMembershipLeaseTime = "CLUSTER_MEMBERSHIP_LEASE_TIME"

	// KeyClusterInformationPathPrefix is the path prefix in the cluster
	// management system (e.g. etcd) under which all keys will be written
	KeyClusterInformationPathPrefix = "CLUSTER_INFORMATION_PATH_PREFIX"

	// KeyBackingServiceEtcd is a whitespace-separated list of host:port for
	// etcd hosts
	KeyBackingServiceEtcd = "BACKINGSERVICE_ETCD"
)
