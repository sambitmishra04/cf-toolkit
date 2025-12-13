package main

type Contest struct {
	ID                  int    `json:"id"`
    Name                string `json:"name"`
    Phase               string `json:"phase"` 
    DurationSeconds     int    `json:"durationSeconds"`
    StartTimeSeconds    int64  `json:"startTimeSeconds"`
}

type ApiResponse struct {
	Status string `json:"status"`
	Result []Contest `json:"result"`
}