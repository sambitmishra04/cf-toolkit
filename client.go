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

	return apiResp.Result, nil
}