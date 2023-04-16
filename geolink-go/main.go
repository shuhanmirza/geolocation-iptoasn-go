package main

import (
	"context"
	"geolink-go/api/service"
	"geolink-go/cronjob"
	"geolink-go/infrastructure"
	"geolink-go/util"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {

	cronScheduler := cron.New()
	go cronScheduler.Run()
	defer cronScheduler.Stop()

	keyValueStore := infrastructure.NewKeyValueStore()

	service.NewGeoDataService(&keyValueStore)

	startCronJobs(cronScheduler, &keyValueStore)

	d, _ := time.ParseDuration("100m")
	time.Sleep(d)
}

func startCronJobs(cronScheduler *cron.Cron, keyValueStore *infrastructure.KeyValueStore) {
	cronJobCtx := context.Background()

	captchaCronJob := cronjob.NewDataBaseSyncCronJob(keyValueStore)
	_, errCron := cronScheduler.AddFunc(util.CronSpecWeekly, func() {
		captchaCronJob.Run(cronJobCtx)
	})

	if errCron != nil {
		log.Println(errCron)
	}
}
