package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	DataRecv   []string `json:"data"`
}

func main() {

	fmt.Println("calling an API")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://jsonmock.hackerrank.com/api/football_matches", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("year", "2011")
	q.Add("team1", "Barcelona")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

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
	fmt.Println(string(data))

	var origin response
	jsonerr := json.Unmarshal(data, &origin)
	fmt.Println(jsonerr)
	fmt.Println("Request Originated from :", origin.Page)

}
