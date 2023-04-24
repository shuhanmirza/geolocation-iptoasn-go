package main

import (
	"context"
	"geolink-go/api/service"
	"geolink-go/api/structs"
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

	geoDataService := service.NewGeoDataService(keyValueStore)

	startCronJobs(cronScheduler, keyValueStore)

	geoDataService.GetIpGeoData(structs.GetGeoDataRequest{
		IpAddress: "1.0.16.1",
	})

	geoDataService.GetIpGeoData(structs.GetGeoDataRequest{
		IpAddress: "221.15.2.2",
	})

	geoDataService.GetIpGeoData(structs.GetGeoDataRequest{
		IpAddress: "221.15.255.255",
	})

	geoDataService.GetIpGeoData(structs.GetGeoDataRequest{
		IpAddress: "103.84.159.230",
	})

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
