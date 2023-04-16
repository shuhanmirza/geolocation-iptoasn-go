package util

import (
	"encoding/binary"
	"encoding/csv"
	"math/big"
	"net"
	"os"
)

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

func ReadCsvFile(csvFileName string) (records [][]string, err error) {
	// open CSV file
	fd, err := os.Open(csvFileName)
	if err != nil {
		return records, err
	}

	defer fd.Close()

	// read CSV file
	fileReader := csv.NewReader(fd)
	return fileReader.ReadAll()
}
