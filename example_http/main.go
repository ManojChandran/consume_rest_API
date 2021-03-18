package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ORIGI_IP string `json:"origin"`
}

func main() {
	fmt.Println("calling an API")
	// For control over HTTP client headers, redirect policy,
	// and other settings, create a Client:
	client := &http.Client{}

	// NewRequest wraps NewRequestWithContext using the background context.
	req, err := http.NewRequest("GET", "https://httpbin.org/ip", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Add custom details to HTTP header request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Do sends an HTTP request and returns an HTTP response,
	// following policy (such as redirects, cookies, auth) as configured on the client.
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	// The client must close the response body when finished with it:
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Convert response body to Response struct
	var origin Response
	json.Unmarshal(data, &origin)
	fmt.Println("Request Originated from :", origin.ORIGI_IP)
}
