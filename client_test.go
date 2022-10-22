package postcodesio_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/leandrorondon/postcodesio-go"
	"github.com/stretchr/testify/assert"
)

var testPostcode = postcodesio.Postcode{
	Postcode:                  "NW1 6XE",
	Outcode:                   "NW1",
	Incode:                    "6XE",
	Quality:                   1,
	Eastings:                  527850,
	Northings:                 182134,
	Country:                   "England",
	NHSHA:                     "London",
	AdminCounty:               "",
	AdminDistrict:             "Westminster",
	AdminWard:                 "Regent's Park",
	Longitude:                 -0.158541,
	Latitude:                  51.523659,
	ParliamentaryConstituency: "Cities of London and Westminster",
	PrimaryCareTrust:          "Westminster",
	Region:                    "London",
	Parish:                    "Westminster, unparished area",
	LSOA:                      "Westminster 008B",
	MSOA:                      "Westminster 008",
	CED:                       "",
	CCG:                       "NHS North West London",
	NUTS:                      "Westminster",
	Codes: postcodesio.Codes{
		AdminCounty:   "E99999999",
		AdminDistrict: "E09000033",
		AdminWard:     "E05013805",
		Parish:        "E43000236",
		CCG:           "E38000256",
		CCGCode:       "",
		NUTS:          "TLI32",
		LAU2:          "E09000033",
		LSOA:          "E01004660",
		MSOA:          "E02000967",
	},
}

type roundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip RoundTripper implementation to be used in tests.
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestNew_WithTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond)
	}))
	defer srv.Close()

	c := postcodesio.NewTestClient(srv.URL, postcodesio.WithTimeout(time.Nanosecond))
	res, err := c.PostcodeLookup(context.Background(), "")

	assert.Nil(t, res)
	assert.ErrorContains(t, err, "context deadline exceeded (Client.Timeout exceeded while awaiting headers)")
}

func TestNew_WithTransport(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status":200}`)
	}))
	defer srv.Close()

	called := false

	var fn roundTripFunc = func(req *http.Request) (*http.Response, error) {
		called = true
		res, _ := http.DefaultTransport.RoundTrip(req)

		return res, nil
	}

	c := postcodesio.NewTestClient(srv.URL, postcodesio.WithTransport(fn))
	res, err := c.PostcodeLookup(context.Background(), "")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, called)
}

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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.RequestURI, "/postcodes")
		assert.Equal(t, http.MethodPost, r.Method)
		body := `{"status":200,"result":[{"query":"NW1 6XE","result":{"postcode":"NW1 6XE","quality":1,"eastings":527850,"northings":182134,"country":"England","nhs_ha":"London","longitude":-0.158541,"latitude":51.523659,"european_electoral_region":"London","primary_care_trust":"Westminster","region":"London","lsoa":"Westminster 008B","msoa":"Westminster 008","incode":"6XE","outcode":"NW1","parliamentary_constituency":"Cities of London and Westminster","admin_district":"Westminster","parish":"Westminster, unparished area","admin_county":null,"admin_ward":"Regent's Park","ced":null,"ccg":"NHS North West London","nuts":"Westminster","codes":{"admin_district":"E09000033","admin_county":"E99999999","admin_ward":"E05013805","parish":"E43000236","parliamentary_constituency":"E14000639","ccg":"E38000256","ccg_id":"W2U3Z","ced":"E99999999","nuts":"TLI32","lsoa":"E01004660","msoa":"E02000967","lau2":"E09000033"}}},{"query":"NW16XE","result":{"postcode":"NW1 6XE","quality":1,"eastings":527850,"northings":182134,"country":"England","nhs_ha":"London","longitude":-0.158541,"latitude":51.523659,"european_electoral_region":"London","primary_care_trust":"Westminster","region":"London","lsoa":"Westminster 008B","msoa":"Westminster 008","incode":"6XE","outcode":"NW1","parliamentary_constituency":"Cities of London and Westminster","admin_district":"Westminster","parish":"Westminster, unparished area","admin_county":null,"admin_ward":"Regent's Park","ced":null,"ccg":"NHS North West London","nuts":"Westminster","codes":{"admin_district":"E09000033","admin_county":"E99999999","admin_ward":"E05013805","parish":"E43000236","parliamentary_constituency":"E14000639","ccg":"E38000256","ccg_id":"W2U3Z","ced":"E99999999","nuts":"TLI32","lsoa":"E01004660","msoa":"E02000967","lau2":"E09000033"}}}]}` //nolint: lll
		fmt.Fprint(w, body)
	}))
	defer srv.Close()

	expected := &postcodesio.BulkPostcodeLookupResponse{
		Status: 200,
		Result: []postcodesio.BulkPostcodeLookupQueryResponse{
			{
				Query:  "NW1 6XE",
				Result: testPostcode,
			},
			{
				Query:  "NW16XE",
				Result: testPostcode},
		},
	}

	c := postcodesio.NewTestClient(srv.URL)
	r, err := c.BulkPostcodeLookup(context.Background(), []string{"NW1 6XE", "NW16XE"})

	assert.NoError(t, err)
	assert.EqualValues(t, expected, r)
}
