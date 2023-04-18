package structs

type GetGeoDataRequest struct {
	IpAddress string `json:"ip_address"`
}

type GetGeoDataResponse struct {
	IpAddress      string `json:"ip_address"`
	CountryCode    string `json:"country_code"`
	Asn            string `json:"asn"`
	AsnDescription string `json:"asn_description"`
}
