package cronjob

import (
	"context"
	"geolink-go/infrastructure"
	"geolink-go/util"
	"log"
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

	cj.downloadDatabase(util.IpToAsnCombinedDatabaseUrl, util.IpToAsnCombinedDatabaseCompressedFilename)
	cj.unzipDatabaseGz(util.IpToAsnCombinedDatabaseCompressedFilename, util.IpToAsnCombinedDatabaseFilename)

	//TODO: Reinitialize the kvstore and binaryList
}

func (cj DatabaseSyncCronJob) downloadDatabase(databaseUrlString string, databaseFilename string) {

	log.Printf("NewDataBaseSyncCronJob: starting database download | url= %s | filename= %s\n", databaseUrlString, databaseFilename)

	written, err := util.DownloadFile(databaseUrlString, databaseFilename)
	if err != nil {
		log.Printf("NewDataBaseSyncCronJob: failed to download %s\n", databaseUrlString)
		log.Println(err)
		return
	}

	log.Printf("NewDataBaseSyncCronJob: database downloaded | url= %s | filename= %s\n | size=%d", databaseUrlString, databaseFilename, written)
}

func (cj DatabaseSyncCronJob) unzipDatabaseGz(databaseCompressedFilename string, databaseFilename string) {
	log.Printf("NewDataBaseSyncCronJob: unzipping database | filename= %s\n", databaseCompressedFilename)

	err := util.UnzipFileGz(databaseCompressedFilename, databaseFilename)
	if err != nil {
		log.Printf("NewDataBaseSyncCronJob: failed to unzip | filename= %s\n", databaseCompressedFilename)
		log.Println(err)
		return
	}

	log.Printf("NewDataBaseSyncCronJob: successfully unzipped database | filename= %s | output=%s\n", databaseCompressedFilename, databaseFilename)
}
