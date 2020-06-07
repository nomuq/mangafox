package scheduler

import (
	"github.com/hibiken/asynq"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func InitilizeCron(queue *asynq.Client) {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/30 * * * *", func() {
		err := Latest("2ZevhabKgkstB6DPzQpMcdSRnxwf78uC", queue)
		if err != nil {
			logrus.Panicln(err)
		}
	})
}
