package util

import (
	"compress/gzip"
	"encoding/binary"
	"encoding/csv"
	"io"
	"math/big"
	"net"
	"net/http"
	"os"
)

func DownloadFile(urlString string, filename string) (written int64, err error) {
	response, err := http.Get(urlString)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// Create the file
	databaseFile, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer databaseFile.Close()

	return io.Copy(databaseFile, response.Body)
}

func UnzipFileGz(compressedFilename string, filename string) (err error) {

	// Open compressed file
	gzipFile, err := os.Open(compressedFilename)
	if err != nil {
		return err
	}

	// Create a gzip reader on top of the file reader
	// Again, it could be any type reader though
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// Uncompress to a writer. We'll use a file writer
	outfileWriter, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	_, err = io.Copy(outfileWriter, gzipReader)
	return err
}

func Ip2Int(ip net.IP) *big.Int {
	i := big.NewInt(0)
	i.SetBytes(ip)
	return i
}

func Int2ip(nn uint32) net.IP {
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

func ReadTsvFile(tsvFileName string) (records [][]string, err error) {
	// open TSV file
	fd, err := os.Open(tsvFileName)
	if err != nil {
		return records, err
	}

	defer fd.Close()

	// read CSV file
	fileReader := csv.NewReader(fd)
	fileReader.Comma = '\t'
	return fileReader.ReadAll()
}
