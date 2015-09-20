package main

import (
	"flag"
	"github.com/forbin"
	log "github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	mesos "github.com/mesos/mesos-go/mesosproto"
	util "github.com/mesos/mesos-go/mesosutil"
	"github.com/mesos/mesos-go/scheduler"
	//"time"
)

const (
	NAME string = "FORBIN"
	CMD  string = "bin/forbin_exec"
)

func main() {
	var master = flag.String("master", "127.0.0.1:5050", "Master address <ip:port>")

	log.Infoln("Lancement PGAgentScheduler") //, time.Now())

	flag.Parse()

	config := scheduler.DriverConfig{
		Master: *master,
		Scheduler: forbin.NewDatabaseScheduler(
			&mesos.ExecutorInfo{
				ExecutorId: util.NewExecutorID("default"),
				Name:       proto.String("FBN"),
				Source:     proto.String("fb_test"),
				Command: &mesos.CommandInfo{
					Value: proto.String(CMD),
					Uris: []*mesos.CommandInfo_URI{
						&mesos.CommandInfo_URI{
							Value:   proto.String("http://localhost:8080/postgresql-bin.tar.gz"),
							Extract: proto.Bool(true),
						},
						&mesos.CommandInfo_URI{
							Value:   proto.String("http://localhost:8080/tools.tar.gz"),
							Extract: proto.Bool(true),
						},
					}},
			},
		),
		Framework: &mesos.FrameworkInfo{
			Name: proto.String(NAME),
			User: proto.String(""),
		},
	}

	driver, err := scheduler.NewMesosSchedulerDriver(config)

	if err != nil {
		log.Fatalln("Unable to create a SchedulerDriver ", err.Error())
	}

	if stat, err := driver.Run(); err != nil {
		log.Infoln("Framework stopped with status %s and error: %s\n", stat.String(), err.Error())
	}

}
