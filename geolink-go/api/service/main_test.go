package service

import (
	"geolink-go/infrastructure"
	"os"
	"testing"
)

var keyValueStore *infrastructure.KeyValueStore
var geoDataService GeoDataService

func TestMain(t *testing.M) {
	os.Chdir("../../")

	keyValueStore = infrastructure.NewKeyValueStore()

	geoDataService = NewGeoDataService(keyValueStore)

	os.Exit(t.Run())
}
