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

	colorClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := colorClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	color := Color{}
	jsonErr := json.Unmarshal(body, &color)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, b := range color.Colors {
		fmt.Printf("%s", b.Value)
	}
}
