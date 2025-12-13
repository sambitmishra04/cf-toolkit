package main

import (
	"net/http"
	"encoding/json"
)

func getContests() ([]Contest, error) {
	resp, err := http.Get("https://codeforces.com/api/contest.list?gym=false")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var apiResp ApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	var upcoming [] Contest

	for _, contest := range apiResp.Result {
		if contest.Phase == "BEFORE"  {
			upcoming = append(upcoming, contest)
		} 
	}

	return upcoming, nil
}