package service

import (
	"geolink-go/api/structs"
	"geolink-go/util"
	"log"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

func TestGeoDataService_GetIpGeoDataV0(t *testing.T) {
	geoDataService.GetIpGeoData(structs.GetGeoDataRequest{
		IpAddress: "1.0.16.1",
	})
}

func TestGeoDataService_GetIpGeoData(t *testing.T) {

	records, _ := util.ReadCsvFile(util.GeoWhoisAsnCountryIpv4NumDatabaseFilename)

	notFoundCount := big.NewInt(0)
	failedCount := big.NewInt(0)
	successfulCount := big.NewInt(0)

	for i, record := range records {
		startIpInt, _ := strconv.Atoi(record[0])
		endIpInt, _ := strconv.Atoi(record[1])
		countryCode := record[2]

		for ipInt := startIpInt; ipInt <= endIpInt; ipInt++ {
			ipAddress := util.Int2ip(uint32(ipInt))
			geoDataResponse, _ := geoDataService.GetIpGeoData(structs.GetGeoDataRequest{IpAddress: ipAddress.String()})
			if geoDataResponse.CountryCode == "None" {
				notFoundCount = notFoundCount.Add(notFoundCount, big.NewInt(1))
			} else if strings.Compare(geoDataResponse.CountryCode, countryCode) != 0 {
				failedCount = failedCount.Add(failedCount, big.NewInt(1))

				//log.Printf("failed | found=%s expected=%s\n", geoDataResponse.CountryCode, countryCode)
			} else {
				successfulCount = successfulCount.Add(successfulCount, big.NewInt(1))
			}
		}
		log.Printf("batch completed | num %d | failed=%d | not found=%d | successful=%d\n", i, failedCount, notFoundCount, successfulCount)
	}

	log.Printf("test done | failed=%d | not found=%d | successful=%d\n", failedCount, notFoundCount, successfulCount)

	//2023/04/20 02:55:03 test done | failed=120964971 | not found=625437929 | successful=2941251342
}
