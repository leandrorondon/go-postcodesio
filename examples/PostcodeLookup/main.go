package main

import (
	"context"
	"fmt"
	"log"

	"github.com/leandrorondon/postcodesio-go"
)

func main() {
	client := postcodesio.New()
	res, err := client.PostcodeLookup(context.Background(), "NW1 6XE")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %d\n", res.Status)
	fmt.Printf("Result: %#v\n", res.Result)
}
