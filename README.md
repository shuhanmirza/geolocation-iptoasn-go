# geolocation-iptoasn-go
**![go-test-badge](https://github.com/shuhanmirza/geolocation-iptoasn-go/actions/workflows/go-test.yml/badge.svg)**
<img src="https://img.shields.io/github/issues/shuhanmirza/geolocation-iptoasn-go.svg?style=flat">

Webservice for getting geolocation and ASN information from IP addresses. This service can be easily run using docker.
It has been implemented using Golang Gin Gonic.

The source of the dataset is https://iptoasn.com/. However, you can easily integrate any database of your liking.

This service loads database into memory and uses binary search to get desired data. It requires around `400MB` of memory.

### Usage

```shell
$ curl 'http://localhost:9000/geo/?ip=103.84.159.230'
```

```json
{
  "ip": "103.84.159.230",
  "country": "BD",
  "asn_number": "133168",
  "asn_description": "SUOSAT-AS-AP Shahjalal University of Science and Technology"
}
```

### Getting Started
#### Option 1:
Prerequisite `go`
```shell
$ cd geolink-go
$ go mod download
$ go build -o main main.go
$ ./main
```

#### Option 2:
Prerequisite `docker`, `docker-compose`
```shell
docker-compose up
```

### Acknowledgement & Inspiration

- https://github.com/jedisct1/iptoasn-webservice
- https://iptoasn.com/
- https://github.com/sapics/ip-location-db
