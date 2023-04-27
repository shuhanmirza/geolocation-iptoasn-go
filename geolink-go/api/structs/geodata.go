package structs

type GetGeoDataRequest struct {
	Ip string `form:"ip" json:"ip" binding:"required"`
}

type GetGeoDataResponse struct {
	Ip             string `json:"ip"`
	Country        string `json:"country"`
	Asn            string `json:"asn"`
	AsnDescription string `json:"asn_description"`
}
