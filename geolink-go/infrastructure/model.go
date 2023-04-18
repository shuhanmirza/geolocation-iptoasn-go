package infrastructure

import "math/big"

type GeoDataModel struct {
	IpRangeStart   *big.Int
	IpRangeEnd     *big.Int
	CountryCode    string
	Asn            string
	AsnDescription string
}
