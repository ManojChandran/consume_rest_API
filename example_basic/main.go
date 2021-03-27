package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Simple get methid
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Print(string(data))
	}

	jsonData := map[string]string{"firstName": "manoj", "scondName": "Chandran"}
	jsonValue, _ := json.Marshal(jsonData)

	// Simple post method
	resp, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Print(err.Error())
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Print(string(data))
	}
}
