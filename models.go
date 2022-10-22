package postcodesio

// Postcode (Ordnance Survey Postcode Directory Dataset).
// Data points returned by the /postcodes and /outcodes API.
type Postcode struct {
	Postcode                  string  `json:"postcode"`
	Outcode                   string  `json:"outcode"`
	Incode                    string  `json:"incode"`
	Quality                   int     `json:"quality"`
	Eastings                  int     `json:"eastings,omitempty"`
	Northings                 int     `json:"northings,omitempty"`
	Country                   string  `json:"country"`
	NHSHA                     string  `json:"nhs_ha,omitempty"` //nolint: tagliatelle
	AdminCounty               string  `json:"admin_county,omitempty"`
	AdminDistrict             string  `json:"admin_district,omitempty"`
	AdminWard                 string  `json:"admin_ward,omitempty"`
	Longitude                 float64 `json:"longitude,omitempty"`
	Latitude                  float64 `json:"latitude,omitempty"`
	ParliamentaryConstituency string  `json:"parliamentary_constituency,omitempty"`
	PrimaryCareTrust          string  `json:"primary_care_trust,omitempty"`
	Region                    string  `json:"region,omitempty"`
	Parish                    string  `json:"parish,omitempty"`
	LSOA                      string  `json:"lsoa,omitempty"`
	MSOA                      string  `json:"msoa,omitempty"`
	CED                       string  `json:"ced,omitempty"`
	CCG                       string  `json:"ccg,omitempty"`
	NUTS                      string  `json:"nuts,omitempty"`
	Codes                     Codes   `json:"codes"`
}

// Codes Represents an ID or Code associated with the postcode.
type Codes struct {
	AdminCounty   string `json:"admin_county,omitempty"`
	AdminDistrict string `json:"admin_district,omitempty"`
	AdminWard     string `json:"admin_ward,omitempty"`
	Parish        string `json:"parish,omitempty"`
	CCG           string `json:"ccg,omitempty"`
	CCGCode       string `json:"ccg_code,omitempty"`
	NUTS          string `json:"nuts,omitempty"`
	LAU2          string `json:"lau2,omitempty"`
	LSOA          string `json:"lsoa,omitempty"`
	MSOA          string `json:"msoa,omitempty"`
}

// Outcode (Ordnance Survey Postcode Directory Dataset).
// Data returned by the /outcodes API.
type Outcode struct {
	Outcode       string   `json:"outcode"`
	Eastings      int      `json:"eastings,omitempty"`
	Northings     int      `json:"northings,omitempty"`
	AdminCounty   []string `json:"admin_county"`
	AdminDistrict []string `json:"admin_district"`
	AdminWard     []string `json:"admin_ward"`
	Longitude     float64  `json:"longitude,omitempty"`
	Latitude      float64  `json:"latitude,omitempty"`
	Country       []string `json:"country"`
	Parish        []string `json:"parish"`
}

// ScottishPostcode (Scottish Postcode Directory).
// Data returned by the /scotland/* APIs.
type ScottishPostcode struct {
	Postcode                          string        `json:"postcode"`
	ScottishParliamentaryConstituency string        `json:"scottish_parliamentary_constituency"`
	Codes                             ScottishCodes `json:"codes"`
}

// ScottishCodes Represents an ID or Code associated with the postcode.
type ScottishCodes struct {
	ScottishParliamentaryConstituency string `json:"scottish_parliamentary_constituency"`
}

// TerminatedPostcode (Ordnance Survey Postcode Directory Dataset).
// Data returned by the /terminated_postcodes/* APIs.
type TerminatedPostcode struct {
	Postcode        string  `json:"postcode"`
	YearTerminated  int     `json:"year_terminated"`
	MonthTerminated int     `json:"month_terminated"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
}

// Place (Ordnance Survey Open Names Dataset).
// Data returned by the /places API.
type Place struct {
	Code                string  `json:"code"`
	Eastings            int     `json:"eastings"`
	Northings           int     `json:"northings"`
	MaxEastings         int     `json:"max_eastings"`
	MinEastings         int     `json:"min_eastings"`
	MaxNorthings        int     `json:"max_northings"`
	MinNorthings        int     `json:"min_northings"`
	Country             string  `json:"country"`
	Longitude           float64 `json:"longitude"`
	Latitude            float64 `json:"latitude"`
	LocalType           string  `json:"local_type"`
	Outcode             string  `json:"outcode"`
	Name1               string  `json:"name1"`
	Name1Lang           string  `json:"name1_lang,omitempty"`
	Name2               string  `json:"name2,omitempty"`
	Name2Lang           string  `json:"name2_lang,omitempty"`
	CountyUnitary       string  `json:"county_unitary,omitempty"`
	CountyUnitaryType   string  `json:"county_unitary_type,omitempty"`
	DistrictBorough     string  `json:"district_borough,omitempty"`
	DistrictBoroughType string  `json:"district_borough_type,omitempty"`
	Region              string  `json:"region"`
}
