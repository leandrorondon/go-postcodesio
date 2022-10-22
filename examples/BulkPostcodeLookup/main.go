package main

import (
	"context"
	"fmt"
	"log"

	"github.com/leandrorondon/postcodesio-go"
)

func main() {
	client := postcodesio.New()

	bulkPostCodeLookup(client)
	bulkPostCodeLookupWithFilters(client)

}

func bulkPostCodeLookup(c *postcodesio.Client) {
	bulkRequest := postcodesio.BulkPostCodeLookupRequest{
		Postcodes: []string{"NW1 6XE", "SW1A 0AA"},
	}
	res, err := c.BulkPostcodeLookup(context.Background(), bulkRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range res.Result {
		fmt.Printf("%s: %#v\n", r.Query, r.Result)
	}
}

func bulkPostCodeLookupWithFilters(c *postcodesio.Client) {
	bulkRequest := postcodesio.BulkPostCodeLookupRequest{
		Postcodes: []string{"NW1 6XE"},
		Filters:   []string{"postcode", "country", "longitude", "latitude"},
	}
	res, err := c.BulkPostcodeLookup(context.Background(), bulkRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %d\n", res.Status)
	for _, r := range res.Result {
		fmt.Printf("%s: %#v\n", r.Query, r.Result)
	}
}
