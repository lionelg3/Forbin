package forbin

import (
	log "github.com/golang/glog"
	exec "github.com/mesos/mesos-go/executor"
	mesos "github.com/mesos/mesos-go/mesosproto"
	run "os/exec"
	"time"
)

type DatabaseExecutor struct {
	tasksLaunched int
	driver        exec.ExecutorDriver
}

func NewDatabaseExecutor() *DatabaseExecutor {
	return &DatabaseExecutor{tasksLaunched: 0}
}

func (self *DatabaseExecutor) Registered(driver exec.ExecutorDriver, execInfo *mesos.ExecutorInfo, fwinfo *mesos.FrameworkInfo, slaveInfo *mesos.SlaveInfo) {
	log.Infoln("Registered Executor on slave ", slaveInfo.GetHostname())
}

func (self *DatabaseExecutor) Reregistered(driver exec.ExecutorDriver, slaveInfo *mesos.SlaveInfo) {
	log.Infoln("Re-registered Executor on slave ", slaveInfo.GetHostname())
}

func (self *DatabaseExecutor) Disconnected(exec.ExecutorDriver) {
	log.Infoln("Executor disconnected.")
}

func (self *DatabaseExecutor) LaunchTask(driver exec.ExecutorDriver, taskInfo *mesos.TaskInfo) {
	log.Infoln("Launching task", taskInfo.GetName(), "with command", taskInfo.Command.GetValue())
	self.driver = driver
	runStatus := &mesos.TaskStatus{
		TaskId: taskInfo.GetTaskId(),
		State:  mesos.TaskState_TASK_RUNNING.Enum(),
	}
	_, err := driver.SendStatusUpdate(runStatus)
	if err != nil {
		log.Infoln("Got error", err)
	}

	self.tasksLaunched++
	log.Infoln("Total tasks launched ", self.tasksLaunched)
	//
	// this is where one would perform the requested task
	//
	self.runProg()
	time.Sleep(time.Second * 40)

	// finish task
	log.Infoln("Finishing task", taskInfo.GetName())
	finStatus := &mesos.TaskStatus{
		TaskId: taskInfo.GetTaskId(),
		State:  mesos.TaskState_TASK_FINISHED.Enum(),
	}
	_, err = driver.SendStatusUpdate(finStatus)
	if err != nil {
		log.Infoln("Got error", err)

	}
	log.Infoln("Task finished", taskInfo.GetName())
}

func (self *DatabaseExecutor) KillTask(exec.ExecutorDriver, *mesos.TaskID) {
	log.Infoln("Kill task")
}

func (self *DatabaseExecutor) FrameworkMessage(driver exec.ExecutorDriver, msg string) {
	log.Infoln("Got framework message: ", msg)
}

func (self *DatabaseExecutor) Shutdown(exec.ExecutorDriver) {
	log.Infoln("Shutting down the executor")
}

func (self *DatabaseExecutor) Error(driver exec.ExecutorDriver, err string) {
	log.Infoln("Got error message:", err)
}

func (self *DatabaseExecutor) runProg() (err error) {
	_, err = run.Command("./pg/init-db.sh").Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = run.Command("./pg/start-db.sh").Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
