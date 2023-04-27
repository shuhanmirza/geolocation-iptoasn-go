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

	ipAddressNet := net.ParseIP(request.Ip)

	if ipAddressNet.IsPrivate() {
		return response, &util.NotPublicIp{}
	}

	ipAddressInt := util.Ip2Int(ipAddressNet)
	startingIpAddress := s.findNearestStartingIp(s.startingIpNumList, ipAddressInt)

	geoDataString := s.keyValueStore.Get(startingIpAddress.String())
	var geoData infrastructure.GeoDataModel
	_ = json.Unmarshal([]byte(geoDataString), &geoData)

	if ipAddressInt.Cmp(geoData.IpRangeEnd) == 1 {
		return response, &util.RecordNotFound{}
	} else if geoData.CountryCode == "None" {
		return response, &util.RecordNotFound{}
	}

	return structs.GetGeoDataResponse{
		Country:        geoData.CountryCode,
		AsnNumber:      geoData.Asn,
		AsnDescription: geoData.AsnDescription,
		Ip:             request.Ip,
	}, nil
}

func (s *GeoDataService) findNearestStartingIp(ipList []*big.Int, candidate *big.Int) (startingIp *big.Int) {

	lenList := len(ipList)

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
