#!/usr/bin/env bash

export GOPATH=$PWD

go fmt github.com/forbin/
go fmt github.com/forbin/repo
go fmt github.com/forbin/scheduler
go fmt github.com/forbin/executor
go build -o bin/ht github.com/forbin/repo
go build -o bin/forbin_sched github.com/forbin/scheduler
go build -o bin/forbin_exec github.com/forbin/executor
