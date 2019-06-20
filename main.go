package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Color struct {
	Colors []struct {
		Value string `json:"value"`
	} `json:"colors"`
}

func main() {
	url := "https://api.noopschallenge.com/hexbot"

	// Create color client
	colorClient := http.Client{
		Timeout: time.Second * 2,
	}

	// Make a GET request to the URL
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Actually, DO the request
	res, getErr := colorClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	// Read the body of the request
	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	// Declare the color struct
	color := Color{}

	// Parses through the json-encoded data and stores it inside &color
	jsonErr := json.Unmarshal(body, &color)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// Iterate through the colors
	for _, b := range color.Colors {
		fmt.Printf("%s", b.Value)
	}
}
