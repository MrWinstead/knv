package cluster

const (
	PathRoot                      = "knv/"
	PathMembers                   = PathRoot + "cluster-members/"
	PathPrefixElections           = PathRoot + "elections/"
	PathPrefixInitializerElection = PathPrefixElections + "cluster-initialization"
	PathClusterInformation        = PathRoot + "cluster-information"
)
