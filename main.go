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

func getColor() string {
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
	var hexValue string
	// Iterate through the colors
	for _, b := range color.Colors {
		// fmt.Printf("%s", b.Value)
		hexValue = b.Value
	}
	return hexValue
}

func handler(w http.ResponseWriter, r *http.Request) {
	color := getColor()
	fmt.Fprintf(w,
		`<body style='background-color: %s'>
			<p style='color: white'>You're looking at hex-value: %s</p>
			<p style='color: white'>Refresh to change the background color.</p>
		</body>`,
		color,
		color,
	)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting up server... check localhost:6060")
	log.Fatal(http.ListenAndServe(":6060", nil))
}
