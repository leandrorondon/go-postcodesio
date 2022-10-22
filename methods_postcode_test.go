package postcodesio_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leandrorondon/postcodesio-go"
	"github.com/stretchr/testify/assert"
)

func TestPostcodeLookup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.RequestURI, "/postcodes/NW1%206XE")
		assert.Equal(t, http.MethodGet, r.Method)
		body := `{"status":200,"result":{"postcode":"NW1 6XE","quality":1,"eastings":527850,"northings":182134,"country":"England","nhs_ha":"London","longitude":-0.158541,"latitude":51.523659,"european_electoral_region":"London","primary_care_trust":"Westminster","region":"London","lsoa":"Westminster 008B","msoa":"Westminster 008","incode":"6XE","outcode":"NW1","parliamentary_constituency":"Cities of London and Westminster","admin_district":"Westminster","parish":"Westminster, unparished area","admin_county":null,"admin_ward":"Regent's Park","ced":null,"ccg":"NHS North West London","nuts":"Westminster","codes":{"admin_district":"E09000033","admin_county":"E99999999","admin_ward":"E05013805","parish":"E43000236","parliamentary_constituency":"E14000639","ccg":"E38000256","ccg_id":"W2U3Z","ced":"E99999999","nuts":"TLI32","lsoa":"E01004660","msoa":"E02000967","lau2":"E09000033"}}}` //nolint: lll
		fmt.Fprint(w, body)
	}))
	defer srv.Close()

	expected := &postcodesio.PostcodeLookupResponse{
		Status: 200,
		Result: testPostcode,
	}

	c := postcodesio.NewTestClient(srv.URL)
	r, err := c.PostcodeLookup(context.Background(), "NW1 6XE")

	assert.NoError(t, err)
	assert.EqualValues(t, expected, r)
}

func TestBulkPostcodeLookup(t *testing.T) {
	tests := []struct {
		name             string
		givenRequest     postcodesio.BulkPostCodeLookupRequest
		responseBody     string
		expectedURL      string
		expectedResponse *postcodesio.BulkPostcodeLookupResponse
	}{
		{
			name: "without filter",
			givenRequest: postcodesio.BulkPostCodeLookupRequest{
				Postcodes: []string{"NW1 6XE", "NW16XE"},
			},
			responseBody: `{"status":200,"result":[{"query":"NW1 6XE","result":{"postcode":"NW1 6XE","quality":1,"eastings":527850,"northings":182134,"country":"England","nhs_ha":"London","longitude":-0.158541,"latitude":51.523659,"european_electoral_region":"London","primary_care_trust":"Westminster","region":"London","lsoa":"Westminster 008B","msoa":"Westminster 008","incode":"6XE","outcode":"NW1","parliamentary_constituency":"Cities of London and Westminster","admin_district":"Westminster","parish":"Westminster, unparished area","admin_county":null,"admin_ward":"Regent's Park","ced":null,"ccg":"NHS North West London","nuts":"Westminster","codes":{"admin_district":"E09000033","admin_county":"E99999999","admin_ward":"E05013805","parish":"E43000236","parliamentary_constituency":"E14000639","ccg":"E38000256","ccg_id":"W2U3Z","ced":"E99999999","nuts":"TLI32","lsoa":"E01004660","msoa":"E02000967","lau2":"E09000033"}}},{"query":"NW16XE","result":{"postcode":"NW1 6XE","quality":1,"eastings":527850,"northings":182134,"country":"England","nhs_ha":"London","longitude":-0.158541,"latitude":51.523659,"european_electoral_region":"London","primary_care_trust":"Westminster","region":"London","lsoa":"Westminster 008B","msoa":"Westminster 008","incode":"6XE","outcode":"NW1","parliamentary_constituency":"Cities of London and Westminster","admin_district":"Westminster","parish":"Westminster, unparished area","admin_county":null,"admin_ward":"Regent's Park","ced":null,"ccg":"NHS North West London","nuts":"Westminster","codes":{"admin_district":"E09000033","admin_county":"E99999999","admin_ward":"E05013805","parish":"E43000236","parliamentary_constituency":"E14000639","ccg":"E38000256","ccg_id":"W2U3Z","ced":"E99999999","nuts":"TLI32","lsoa":"E01004660","msoa":"E02000967","lau2":"E09000033"}}}]}`, //nolint: lll
			expectedURL:  "/postcodes",
			expectedResponse: &postcodesio.BulkPostcodeLookupResponse{
				Status: 200,
				Result: []postcodesio.BulkPostcodeLookupQueryResponse{
					{Query: "NW1 6XE", Result: testPostcode},
					{Query: "NW16XE", Result: testPostcode}},
			},
		},
		{
			name: "with filter",
			givenRequest: postcodesio.BulkPostCodeLookupRequest{
				Postcodes: []string{"NW1 6XE"},
				Filters:   []string{"postcode", "country", "longitude", "latitude"},
			},
			responseBody: `{"status":200,"result":[{"query":"NW1 6XE","result":{"postcode":"NW1 6XE","country":"England","longitude":-0.158541,"latitude":51.523659}}]}`, //nolint: lll
			expectedURL:  "/postcodes?filter=postcode,country,longitude,latitude",
			expectedResponse: &postcodesio.BulkPostcodeLookupResponse{
				Status: 200,
				Result: []postcodesio.BulkPostcodeLookupQueryResponse{
					{
						Query: "NW1 6XE",
						Result: postcodesio.Postcode{
							Postcode:  "NW1 6XE",
							Country:   "England",
							Longitude: -0.158541,
							Latitude:  51.523659,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.expectedURL, r.RequestURI)
				assert.Equal(t, http.MethodPost, r.Method)
				fmt.Fprint(w, test.responseBody)
			}))
			defer srv.Close()

			c := postcodesio.NewTestClient(srv.URL)
			r, err := c.BulkPostcodeLookup(context.Background(), test.givenRequest)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expectedResponse, r)
		})
	}
}
