package service

import (
	"geolink-go/api/structs"
	"testing"
)

func TestGeoDataService_GetIpGeoData(t *testing.T) {
	ipAddress := "103.84.159.230"
	ipLocation := "BD"

	data, err := geoDataService.GetIpGeoData(structs.GetGeoDataRequest{
		Ip: ipAddress,
	})
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}
	if data.Country != ipLocation {
		t.Errorf(" location of %s is %s but found %s", ipAddress, ipLocation, data.Country)
	}
}

func TestGeoDataService_GetIpGeoDataWithExistingDataset(t *testing.T) {

	/*	commented out on purpose. uncomment when needed. this test requires more than 3 hours to complete

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
				geoDataResponse, _ := geoDataService.GetIpGeoData(structs.GetGeoDataRequest{Ip: ipAddress.String()})
				if geoDataResponse.Country == "None" {
					notFoundCount = notFoundCount.Add(notFoundCount, big.NewInt(1))
				} else if strings.Compare(geoDataResponse.Country, countryCode) != 0 {
					failedCount = failedCount.Add(failedCount, big.NewInt(1))

					//log.Printf("failed | found=%s expected=%s\n", geoDataResponse.Country, countryCode)
				} else {
					successfulCount = successfulCount.Add(successfulCount, big.NewInt(1))
				}
			}
			log.Printf("batch completed | num %d | failed=%d | not found=%d | successful=%d\n", i, failedCount, notFoundCount, successfulCount)
		}

		log.Printf("test done | failed=%d | not found=%d | successful=%d\n", failedCount, notFoundCount, successfulCount)

		//2023/04/20 02:55:03 test done | failed=120964971 | not found=625437929 | successful=2941251342


	*/
}
