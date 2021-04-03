package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// https://mholt.github.io/json-to-go/
type response struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		Competition string `json:"competition"`
		Year        int    `json:"year"`
		Round       string `json:"round"`
		Team1       string `json:"team1"`
		Team2       string `json:"team2"`
		Team1Goals  string `json:"team1goals"`
		Team2Goals  string `json:"team2goals"`
	} `json:"data"`
}

type request struct {
	endpoint    string
	team        string
	year        string
	matchplayed string
	page        string
}

func (r *request) teamData() response {
	var origin response
	client := &http.Client{}
	req, err := http.NewRequest("GET", r.endpoint, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Add custom details to HTTP header request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("year", r.year)
	if r.matchplayed == "away" {
		q.Add("team2", r.team)
	} else {
		q.Add("team1", r.team)
	}
	q.Add("page", r.page)

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	// Do sends an HTTP request and returns an HTTP response,
	// following policy (such as redirects, cookies, auth) as configured on the client.
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	json.Unmarshal(data, &origin)
	return origin
}

func (res *response) calculateGoals() int32 {
	var TotalGoals, T int
	fmt.Println(res.Data)
	for _, c := range res.Data {
		//fmt.Println("goals", c.Team1Goals)
		T, _ = strconv.Atoi(c.Team1Goals)
		TotalGoals = TotalGoals + T
	}
	return int32(TotalGoals)
}

func main() {
	fmt.Println("calling an API")
	var r request
	var awaygoals, homegoals int32
	var i int32
	var homerecord, awayrecord response
	//"https://jsonmock.hackerrank.com/api/football_competition"
	r.endpoint = "https://jsonmock.hackerrank.com/api/football_matches"
	r.team = "Barcelona"
	r.year = "2011"
	r.matchplayed = "home"

	for i = 1; len(homerecord.Data) == 0; i++ {
		r.page = strconv.Itoa(int(i))
		homerecord = r.teamData()
		homegoals = homegoals + homerecord.calculateGoals()
		//fmt.Println(homerecord.calculateGoals())
	}

	r.matchplayed = "away"

	for i = 1; len(awayrecord.Data) == 0; i++ {
		r.page = strconv.Itoa(int(i))
		awayrecord = r.teamData()
		awaygoals = awaygoals + awayrecord.calculateGoals()
		//fmt.Println(awayrecord.calculateGoals())
	}

	fmt.Println(awaygoals, "+", homegoals, "=", awaygoals+homegoals)
}
