package cron

import (
	"github.com/robfig/cron/v3"
	"time"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	Cron.AddFunc("0 0 0 * * *", SyncArticleData)
	Cron.AddFunc("0 0 0 * * *", SyncCommentData)
	Cron.Start()
}
