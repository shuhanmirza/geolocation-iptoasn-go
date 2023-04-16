package cronjob

import (
	"context"
	"geolink-go/infrastructure"
	"geolink-go/util"
	"io"
	"log"
	"net/http"
	"os"
)

type DatabaseSyncCronJob struct {
	keyValueStore *infrastructure.KeyValueStore
}

func NewDataBaseSyncCronJob(keyValueStore *infrastructure.KeyValueStore) DatabaseSyncCronJob {
	return DatabaseSyncCronJob{keyValueStore: keyValueStore}
}

func (cj DatabaseSyncCronJob) Run(_ context.Context) {
	cj.downloadDatabase(util.GeoWhoisAsnCountryIpv4NumDatabaseUrl, util.GeoWhoisAsnCountryIpv4NumDatabaseFilename)
	cj.downloadDatabase(util.GeoWhoisAsnCountryIpv6NumDatabaseUrl, util.GeoWhoisAsnCountryIpv6NumDatabaseFilename)
}

func (cj DatabaseSyncCronJob) downloadDatabase(databaseUrlString string, databaseFilename string) {

	log.Printf("NewDataBaseSyncCronJob: starting database download | url= %s | filename= %s\n", databaseUrlString, databaseFilename)

	response, err := http.Get(databaseUrlString)
	if err != nil {
		log.Printf("NewDataBaseSyncCronJob: failed to download %s\n", databaseUrlString)
		log.Println(err)
		return
	}
	defer response.Body.Close()

	// Create the file
	databaseFile, err := os.Create(databaseFilename)
	if err != nil {
		log.Printf("NewDataBaseSyncCronJob: failed to create file %s\n", databaseFilename)
		log.Println(err)
		return
	}
	defer databaseFile.Close()

	written, err := io.Copy(databaseFile, response.Body)
	if err != nil {
		log.Printf("NewDataBaseSyncCronJob: Error while populating the file %s\n", databaseFilename)
		log.Println(err)
		return
	}

	log.Printf("NewDataBaseSyncCronJob: database downloaded | url= %s | filename= %s\n | size=%d", databaseUrlString, databaseFilename, written)
}
