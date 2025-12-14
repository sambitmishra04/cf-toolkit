package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func main() { 

	initDB()

	fmt.Println("Initializing Google Calendar...")
    srv := getCalendarService()
    fmt.Println("Success! Authenticated.")

	runSync(srv)

	ticker := time.NewTicker(24 * time.Hour)
	fmt.Println("Worker started. Checking every 24 hours...")

	for range ticker.C {
		runSync(srv)
	}
}
func runSync(srv *calendar.Service) {
	fmt.Println("Checking for new contests...")
	contests, err := getContests()
	if err != nil {
		log.Printf("Error fetching contests: %v", err)
		return
	}

	for _, c := range contests {
		if eventExists(c.ID) {
			fmt.Printf("Skipping %s (already added)\n", c.Name)
			continue
		}

		t := time.Unix(c.StartTimeSeconds, 0)
		fmt.Printf("- Adding %s\n  When: %s\n", c.Name, t.Format(time.RFC1123))

		addContestToCalendar(srv, c)
		saveEvent(c.ID)
	}
}