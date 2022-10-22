package postcodesio

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// PostcodeLookup This uniquely identifies a postcode.
// Returns a single postcode entity for a given postcode (case, space insensitive).
// If no postcode is found it returns "404" response code.
// GET https://api.postcodes.io/postcodes/:postcode
func (c *Client) PostcodeLookup(ctx context.Context, postcode string) (*PostcodeLookupResponse, error) {
	url := fmt.Sprintf("%s/postcodes/%s", c.baseURL, postcode)

	b, err := c.get(ctx, url)
	if err != nil {
		return nil, err
	}

	var r PostcodeLookupResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// BulkPostcodeLookup Accepts a JSON object containing an array of postcodes. Returns a list of matching postcodes and
// respective available data. Accepts up to 100 postcodes.
// POST https://api.postcodes.io/postcodes
func (c *Client) BulkPostcodeLookup(ctx context.Context, bulkRequest BulkPostCodeLookupRequest) (*BulkPostcodeLookupResponse, error) {
	url := fmt.Sprintf("%s/postcodes", c.baseURL)

	if len(bulkRequest.Filters) > 0 {
		filters := strings.Join(bulkRequest.Filters, ",")
		url = fmt.Sprintf("%s?filter=%s", url, filters)
	}

	b, err := c.post(ctx, url, bulkRequest)
	if err != nil {
		return nil, err
	}

	var r BulkPostcodeLookupResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// ReverseGeocoding Returns nearest postcodes for a given longitude and latitude.
// GET https://api.postcodes.io/postcodes?lon=:longitude&lat=:latitude
func (c *Client) ReverseGeocoding(ctx context.Context, request ReverseGeocodingRequest) (*ReverseGeocodingResponse, error) {
	url := fmt.Sprintf("%s/postcodes?lon=%g&lat=%g", c.baseURL, request.Longitude, request.Latitude)

	if request.Limit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, request.Limit)
	}

	if request.Radius > 0 {
		url = fmt.Sprintf("%s&radius=%g", url, request.Radius)
	}

	if request.WideSearch {
		url = fmt.Sprintf("%s&widesearch=true", url)
	}

	b, err := c.get(ctx, url)
	if err != nil {
		return nil, err
	}

	var r ReverseGeocodingResponse
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
