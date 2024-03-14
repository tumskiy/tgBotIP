package types

type Site struct {
	IP         string  `json:"ip"`
	IPDecimal  int     `json:"ip_decimal"`
	Country    string  `json:"country"`
	CountryIso string  `json:"country_iso"`
	CountryEu  bool    `json:"country_eu"`
	RegionName string  `json:"region_name"`
	RegionCode string  `json:"region_code"`
	ZipCode    string  `json:"zip_code"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	TimeZone   string  `json:"time_zone"`
	Asn        string  `json:"asn"`
	AsnOrg     string  `json:"asn_org"`
	UserAgent  struct {
		Product  string `json:"product"`
		Version  string `json:"version"`
		RawValue string `json:"raw_value"`
	} `json:"user_agent"`
}
