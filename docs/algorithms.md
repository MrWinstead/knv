# KNV Algorithms

## Cluster Membership

Nodes are expected to maintain membership within the cluster. Using etcd as a
central storage location for such information, cluster members register their
server identity under a key prefix. While running, each member of the cluster
continuously maintains the lease to which their cluster membership information
is bound.

Consider a cluster with the following Nodes:
* 0 -> `01D5X95JXK9SZN5PHK2YVGPD5Z`
* 1 -> `01D5X95T0F171R57TCT8F6KKQB`

Etcd will have the following directory structure:
```
$ etcdctl get --prefix /github.com/mrwinstead/knv/knv/cluster-members/
/github.com/mrwinstead/knv/knv/cluster-members/01D5X95JXK9SZN5PHK2YVGPD5Z
/github.com/mrwinstead/knv/knv/cluster-members/01D5X95T0F171R57TCT8F6KKQB
```

### Node Failure

Each node's membership information is stored assigned to a lease. When the lease
expires, etcd will delete the path, and the node is considered offline.
