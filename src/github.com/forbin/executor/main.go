package main

import (
	"github.com/forbin"
	log "github.com/golang/glog"
	exec "github.com/mesos/mesos-go/executor"
)

func main() {
	log.Infoln("START executor")

	config := exec.DriverConfig{
		Executor: forbin.NewDatabaseExecutor(),
	}

	driver, err := exec.NewMesosExecutorDriver(config)

	if err != nil {
		log.Infoln("Unable to create a ExecutorDriver ", err.Error())
	}

	_, err = driver.Start()
	if err != nil {
		log.Infoln("Got error:", err)
		return
	}
	log.Infoln("Executor process has started and running.")
	driver.Join()
}
