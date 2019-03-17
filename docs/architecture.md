# Architecture

# Cluster Topology

A cluster is described as a series of nodes which vie for write leadership over
a number of shards of the keyspace. The keyspace is divided upon cluster
declaration into a predefined number of shards.

![cluster topology](./cluster%20topology.png)

Notes:
* Nodes calculate their nodeID integer value by fetching the identities of all
cluster members, lexographically sorting them (since node identities are 
[ulids](https://github.com/ulid/spec)), then ascertaining the offset into array
of IDs of their ID.

