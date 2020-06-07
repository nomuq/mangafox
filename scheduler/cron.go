package scheduler

import (
	"github.com/hibiken/asynq"
)

func InitilizeCron(queue *asynq.Client) {
	// c := cron.New(cron.WithSeconds())

	// c.AddFunc("*/30 * * * *", func() {
	// 	err := Latest("2ZevhabKgkstB6DPzQpMcdSRnxwf78uC", queue)
	// 	if err != nil {
	// 		logrus.Errorln(err)
	// 	}
	// })

	// c.AddFunc("* * * * *", func() {
	// 	task := asynq.NewTask(string(tasks.UpdateSearchIndexes), map[string]interface{}{"type": "manga"})
	// 	err := queue.Enqueue(task, asynq.Unique(time.Hour), asynq.MaxRetry(0))
	// 	if err != nil {
	// 		logrus.Errorln(err)
	// 	}
	// })

}
