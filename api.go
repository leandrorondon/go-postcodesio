package postcodesio

// Postcodes

// PostcodeLookupResponse represents the response of the Postcode Lookup API method.
type PostcodeLookupResponse struct {
	Status int      `json:"status"`
	Result Postcode `json:"result"`
}

// BulkPostCodeLookupRequest is the input for the Bulk Postcode Lookup API method.
type BulkPostCodeLookupRequest struct {
	Postcodes []string `json:"postcodes"`
	Filters   []string `json:"-"`
}

// BulkPostcodeLookupResponse represents the response of the Bulk Postcode Lookup API method.
type BulkPostcodeLookupResponse struct {
	Status int                               `json:"status"`
	Result []BulkPostcodeLookupQueryResponse `json:"result"`
}

// BulkPostcodeLookupQueryResponse is the result of a query from Bulk Postcode Lookup response.
type BulkPostcodeLookupQueryResponse struct {
	Query  string   `json:"query"`
	Result Postcode `json:"result"`
}
