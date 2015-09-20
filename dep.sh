#!/usr/bin/env bash

export GOPATH=$PWD

# Mesos
go get github.com/mesos/mesos-go
go get github.com/pborman/uuid
go get github.com/samuel/go-zookeeper/zk
go get golang.org/x/net/context
go get github.com/golang/glog
go get github.com/stretchr/testify/mock
go get github.com/gogo/protobuf/proto
go get github.com/golang/protobuf/proto

# Postgresql
go get github.com/lib/pq
