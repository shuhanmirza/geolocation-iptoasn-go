package structs

type GetGeoDataRequest struct {
	Ip string `form:"ip" json:"ip" binding:"required"`
}

type GetGeoDataResponse struct {
	Ip             string `json:"ip"`
	Country        string `json:"country"`
	AsnNumber      string `json:"asn_number"`
	AsnDescription string `json:"asn_description"`
}
