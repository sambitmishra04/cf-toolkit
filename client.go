package main

import (
	"io"
	"log"
	"net/http"
)

func getContests() string {
	resp, err := http.Get("https://codeforces.com/api/contest.list?gym=false")
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	return string(body)
}