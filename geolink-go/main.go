package main

import (
	"context"
	"geolink-go/api/controller"
	"geolink-go/api/routes"
	"geolink-go/api/service"
	"geolink-go/cronjob"
	"geolink-go/infrastructure"
	"geolink-go/util"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {

	cronScheduler := cron.New()
	go cronScheduler.Run()
	defer cronScheduler.Stop()

	keyValueStore := infrastructure.NewKeyValueStore()

	// initialize gin router
	log.Println("Initializing Routes")
	ginRouter := infrastructure.NewGinRouter()

	geoDataService := service.NewGeoDataService(keyValueStore)
	geoDataController := controller.NewGeoDataController(geoDataService)
	geoDataRoute := routes.NewGeoDataRoute(geoDataController, ginRouter)
	geoDataRoute.Setup()

	//server
	serverAddress := util.ServerAddress
	err := ginRouter.Gin.Run(serverAddress)
	if err != nil {
		log.Println(err)
		log.Fatal("could not start APIs")
	}

	log.Println("server started")

	startCronJobs(cronScheduler, keyValueStore)
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
