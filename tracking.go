package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetTrackingShipment(apiKeyValue string, trackingNumber *string) {
	url := "https://api-eu.dhl.com/track/shipments?trackingNumber=" + *trackingNumber
	fmt.Println("GET:", url)

	client := &http.Client{} // Create a new HTTP client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Create a new GET request was not possible due to:", err)
		os.Exit(1)
	}

	req.Header.Add("DHL-API-Key", apiKeyValue) // Add DHL API Key to header GET request

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Send the GET request was not possible due to:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Unexpected HTTP response status code:", res.StatusCode)
		os.Exit(1) // Throwing an error when HTTP status code is != 200
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read the response body was not possible due to:", err)
		os.Exit(1)
	}

	var jsonData any
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	responseBodyJSONformat, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON for printing:", err)
		os.Exit(1)
	}

	fmt.Println(string(responseBodyJSONformat))
}
