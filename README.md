# KNV, a Distributed Multi-Key Value Database

## Overview
The KNV project aims to be a learning tool for a basic key-value database with
two twists:
* A single value may be listed under multiple indexes within a table
* Distributed Database

## [Documentation](./docs/index.md)

## Deployment Requirements

KNV expects to run with the following services:
* Etcd

## Roadmap

The following are currently supported:
* gRPC-based multi-key-value store API
* local on-disk storage
* Cluster Membership

Up next are the following features:
* Write re-direction to shard write-leaders
* Kafka-backed storage
* Kafka-enabled storage replication
* Separate read instances
* Self-healing node/shard assignment
