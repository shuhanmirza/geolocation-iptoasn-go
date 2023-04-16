package service

import (
	"geolink-go/infrastructure"
	"geolink-go/util"
	"log"
)

type GeoDataService struct {
	keyValueStore *infrastructure.KeyValueStore
}

func NewGeoDataService(keyValueStore *infrastructure.KeyValueStore) GeoDataService {
	service := GeoDataService{
		keyValueStore: keyValueStore,
	}

	go service.loadKeyValueStore()

	return service
}

func (s GeoDataService) loadKeyValueStore() {
	log.Println("GeoDataService: starting loading csv file")

	records, err := util.ReadCsvFile(util.GeoWhoisAsnCountryIpv4NumDatabaseFilename)
	if err != nil {
		log.Println("GeoDataService: error while loading csv file")
		log.Println(err)
		return
	}

	log.Println("GeoDataService: csv file loaded successfully")

	log.Println(len(records))
}
