package forbin

import (
	log "github.com/golang/glog"
	proto "github.com/golang/protobuf/proto"
	. "github.com/mesos/mesos-go/mesosproto"
	mesos "github.com/mesos/mesos-go/mesosproto"
	. "github.com/mesos/mesos-go/mesosutil"
	util "github.com/mesos/mesos-go/mesosutil"
	. "github.com/mesos/mesos-go/scheduler"
	"strconv"
)

type DatabaseScheduler struct {
	name            string
	taskOid         int
	databaseRunning bool
	execInfo        *ExecutorInfo
}

func NewDatabaseScheduler(executor *ExecutorInfo) *DatabaseScheduler {
	return &DatabaseScheduler{
		taskOid:  0,
		execInfo: executor,
	}
}

func (scheduler *DatabaseScheduler) Registered(driver SchedulerDriver, frameworkId *FrameworkID, masterInfo *MasterInfo) {
	log.Infoln("Framework Registered with Master ", masterInfo)
}

func (scheduler *DatabaseScheduler) Reregistered(driver SchedulerDriver, masterInfo *MasterInfo) {
	log.Infoln("Framework Re-Registered with Master ", masterInfo)
}

func (scheduler *DatabaseScheduler) ResourceOffers(driver SchedulerDriver, offers []*Offer) {
	for _, offer := range offers {
		cpuResources := FilterResources(offer.Resources, func(res *Resource) bool {
			return res.GetName() == "cpus"
		})
		cpus := 0.0
		for _, res := range cpuResources {
			cpus += res.GetScalar().GetValue()
		}

		memResources := FilterResources(offer.Resources, func(res *Resource) bool {
			return res.GetName() == "mem"
		})
		mems := 0.0
		for _, res := range memResources {
			mems += res.GetScalar().GetValue()
		}

		log.Infoln("Received Offer <", offer.Id.GetValue(), "> with cpus =", cpus, " mem =", mems)

		scheduler.taskOid++

		taskId := &TaskID{
			Value: proto.String(strconv.Itoa(scheduler.taskOid)),
		}

		if scheduler.databaseRunning {
			return
		}

		var tasks []*mesos.TaskInfo
		task := &mesos.TaskInfo{
			Name:     proto.String("forbin-task-" + taskId.GetValue()),
			TaskId:   taskId,
			SlaveId:  offer.SlaveId,
			Executor: scheduler.execInfo,
			Resources: []*mesos.Resource{
				util.NewScalarResource("cpus", 1),
				util.NewScalarResource("mem", 128),
			},
		}
		tasks = append(tasks, task)
		log.Infoln("Launching ", len(tasks), "tasks for offer", offer.Id.GetValue())
		driver.LaunchTasks([]*mesos.OfferID{offer.Id}, tasks, &mesos.Filters{RefuseSeconds: proto.Float64(1)})
		scheduler.databaseRunning = true
	}
}

func (scheduler *DatabaseScheduler) Disconnected(SchedulerDriver) {
	log.Infoln("Disconnected")
}

func (scheduler *DatabaseScheduler) StatusUpdate(driver SchedulerDriver, status *TaskStatus) {
	log.Infoln("StatusUpdate", status)
	if *status.State != TaskState_TASK_RUNNING {
		scheduler.databaseRunning = false
	}
}

func (scheduler *DatabaseScheduler) OfferRescinded(SchedulerDriver, *OfferID) {
	log.Infoln("OfferRescinded")
}

func (scheduler *DatabaseScheduler) FrameworkMessage(driver SchedulerDriver, execId *ExecutorID, slaveId *SlaveID,
	data string) {
	log.Infoln("FrameworkMessage " + data)
}

func (scheduler *DatabaseScheduler) SlaveLost(SchedulerDriver, *SlaveID) {
	log.Infoln("SlaveLost")
}

func (scheduler *DatabaseScheduler) ExecutorLost(SchedulerDriver, *ExecutorID, *SlaveID, int) {
	log.Infoln("Launching")
}

func (scheduler *DatabaseScheduler) Error(driver SchedulerDriver, err string) {
	log.Infoln("PostgresqlScheduler received error:", err)
}
