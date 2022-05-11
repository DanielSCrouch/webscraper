package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	uri := "https://en.wikipedia.org/wiki/Ken_Thompson"
	log.Printf("fetching: %s", uri)

	response, err := http.Get(uri)
	if err != nil {
		log.Fatalf("error retrieving site `%s`: %s", uri, err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error reading response body: %s", err)
	}

	log.Printf("%d bytes received", len(body))
	log.Printf("page body (first 1000 bytes):\n%s...", string(body[0:1000]))
}
