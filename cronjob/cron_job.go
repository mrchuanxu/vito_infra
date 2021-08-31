package cronjob

import (
	"github.com/vito_infra/logger"
	"github.com/robfig/cron"
)

// CronJob cronjob job
type CronJob struct {
	Name string
	Spec string
	Job  func()
}

func RunCron(c *cron.Cron, cronJobs []*CronJob) {
	if len(cronJobs) == 0 {
		return
	}
	for i := range cronJobs {
		logger.TransLogger.Sugar().Infof("cronjob run jobs:[%s],spec:[%s]", cronJobs[i].Name, cronJobs[i].Spec)
		c.AddFunc(cronJobs[i].Spec, cronJobs[i].Job)
	}
	c.Start()
}
