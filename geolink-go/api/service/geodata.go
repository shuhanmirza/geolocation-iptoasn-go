package service

import (
	"encoding/json"
	"geolink-go/api/structs"
	"geolink-go/infrastructure"
	"geolink-go/util"
	"log"
	"math/big"
	"net"
	"sort"
)

type GeoDataService struct {
	keyValueStore     *infrastructure.KeyValueStore
	startingIpNumList []*big.Int
}

func NewGeoDataService(keyValueStore *infrastructure.KeyValueStore) GeoDataService {
	service := GeoDataService{
		keyValueStore: keyValueStore,
	}

	service.loadKeyValueStore()

	return service
}

func (s *GeoDataService) GetIpGeoData(request structs.GetGeoDataRequest) (response structs.GetGeoDataResponse, err error) {

	//startTime := time.Now().Nanosecond()

	ipAddressInt := util.Ip2Int(net.ParseIP(request.IpAddress))

	//log.Printf("requested ip=%s | ipint=%d\n", request.IpAddress, ipAddressInt)

	startingIpAddress := s.findNearestStartingIp(s.startingIpNumList, ipAddressInt)

	geoDataString := s.keyValueStore.Get(startingIpAddress.String())

	//log.Println(geoDataString)

	var geoData infrastructure.GeoDataModel
	_ = json.Unmarshal([]byte(geoDataString), &geoData)

	//log.Printf("ip range %s - %s\n", util.Int2ip(uint32(geoData.IpRangeStart.Int64())), util.Int2ip(uint32(geoData.IpRangeEnd.Int64())))

	//log.Printf("time required %d\n", time.Now().Nanosecond()-startTime)

	return structs.GetGeoDataResponse{
		CountryCode:    geoData.CountryCode,
		Asn:            geoData.Asn,
		AsnDescription: geoData.AsnDescription,
		IpAddress:      request.IpAddress,
	}, nil
}

func (s *GeoDataService) findNearestStartingIp(ipList []*big.Int, candidate *big.Int) (startingIp *big.Int) {

	lenList := len(ipList)

	//if lenList < 10 {
	//	log.Println(ipList)
	//}

	if lenList == 1 {
		return ipList[0]
	} else if lenList == 2 {
		if candidate.Cmp(ipList[1]) >= 0 {
			return ipList[1]
		}
		return ipList[0]
	}

	if ipList[lenList/2].Cmp(candidate) == 0 {
		return ipList[lenList/2]
	} else if ipList[lenList/2].Cmp(candidate) == -1 {
		return s.findNearestStartingIp(ipList[lenList/2:lenList], candidate)
	} else {
		return s.findNearestStartingIp(ipList[0:lenList/2], candidate)
	}

}

func (s *GeoDataService) loadKeyValueStore() {
	log.Println("GeoDataService: starting loading database files")

	records, err := util.ReadTsvFile(util.IpToAsnCombinedDatabaseFilename)
	if err != nil {
		log.Printf("GeoDataService: error while loading %s\n", util.IpToAsnCombinedDatabaseFilename)
		log.Println(err)
		return
	}

	for _, record := range records {
		startingIpString := record[0]
		endingIpString := record[1]
		asnString := record[2]
		countryCodeString := record[3]
		asnDescriptionString := record[4]

		startingIpInt := util.Ip2Int(net.ParseIP(startingIpString))
		endingIpInt := util.Ip2Int(net.ParseIP(endingIpString))

		s.startingIpNumList = append(s.startingIpNumList, startingIpInt)

		data := infrastructure.GeoDataModel{
			IpRangeStart:   startingIpInt,
			IpRangeEnd:     endingIpInt,
			Asn:            asnString,
			AsnDescription: asnDescriptionString,
			CountryCode:    countryCodeString,
		}
		dataString, _ := json.Marshal(data)

		s.keyValueStore.Set(startingIpInt.String(), string(dataString))
	}

	sort.Slice(s.startingIpNumList, func(i, j int) bool {
		return s.startingIpNumList[i].Cmp(s.startingIpNumList[j]) < 0
	})

	log.Printf("GeoDataService: database loading completed | number of item: %d\n", len(s.startingIpNumList))
}
