package linda

import (
	"github.com/amlun/linda/linda/core"
	"github.com/twinj/uuid"
)

func (l *Linda) Schedule() (string, error) {
	Logger.Info("schedule task from smarter")
	return l.smarter.GetTask()
}

func (l *Linda) ScheduleTask(task *core.Task) func() {
	return func() {
		var job core.Job
		job.JobId = uuid.NewV4().String()
		job.TaskId = task.TaskId
		job.Queue = task.Queue
		job.Payload = task.Payload
		l.PushJob(job)
	}
}
