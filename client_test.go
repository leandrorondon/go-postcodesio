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
