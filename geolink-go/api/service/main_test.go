package service

import (
	"geolink-go/infrastructure"
	"os"
	"testing"
)

var keyValueStore *infrastructure.KeyValueStore
var geoDataService GeoDataService

func TestMain(t *testing.M) {
	keyValueStore = infrastructure.NewKeyValueStore()

	geoDataService = NewGeoDataService(keyValueStore)

	os.Exit(t.Run())

	//Note: you need to change paths of the csv files in constant to run the tests. I am too lazy. Sorry.
}
