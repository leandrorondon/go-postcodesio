package main

import (
	"context"
	"fmt"
	"log"

	"github.com/leandrorondon/postcodesio-go"
)

func main() {
	client := postcodesio.New()

	reverseGeocoding(client)
	reverseGeocodingWithExtras(client)

}

func reverseGeocoding(c *postcodesio.Client) {
	request := postcodesio.ReverseGeocodingRequest{
		Longitude: -0.158541,
		Latitude:  51.523659,
	}
	res, err := c.ReverseGeocoding(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range res.Result {
		fmt.Printf("%#v\n", r)
	}
}

func reverseGeocodingWithExtras(c *postcodesio.Client) {
	request := postcodesio.ReverseGeocodingRequest{
		Longitude:  -0.158541,
		Latitude:   51.5236,
		Radius:     6.6,
		Limit:      1,
		WideSearch: true,
	}
	res, err := c.ReverseGeocoding(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range res.Result {
		fmt.Printf("%#v\n", r)
	}
}
