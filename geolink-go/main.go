package main

import (
	"context"
	"encoding/binary"
	"geolink-go/cronjob"
	"geolink-go/infrastructure"
	"geolink-go/util"
	"github.com/robfig/cron/v3"
	"log"
	"math/big"
	"net"
	"time"
)

func main() {

	cronScheduler := cron.New()
	go cronScheduler.Run()
	defer cronScheduler.Stop()

	keyValueStore := infrastructure.NewKeyValueStore()

	startCronJobs(cronScheduler, &keyValueStore)

	d, _ := time.ParseDuration("100m")
	time.Sleep(d)
}

func startCronJobs(cronScheduler *cron.Cron, keyValueStore *infrastructure.KeyValueStore) {
	cronJobCtx := context.Background()

	captchaCronJob := cronjob.NewDataBaseSyncCronJob(keyValueStore)
	_, errCron := cronScheduler.AddFunc(util.CronSpecEveryOneMin, func() {
		captchaCronJob.Run(cronJobCtx)
	})

	if errCron != nil {
		log.Println(errCron)
	}
}

func ipv6ToInt(IPv6Addr net.IP) *big.Int {
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(IPv6Addr)
	return IPv6Int
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		panic("no sane way to convert ipv6 into uint32")
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
