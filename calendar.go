package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func addContestToCalendar(srv *calendar.Service, contest Contest) {
	start := time.Unix(contest.StartTimeSeconds, 0)
	end := start.Add(time.Duration(contest.DurationSeconds) * time.Second)

	event := &calendar.Event{
		Summary:     contest.Name,
		Location:    "Codeforces",
		Description: fmt.Sprintf("https://codeforces.com/contest/%d", contest.ID),
		Start: &calendar.EventDateTime{
			DateTime: start.Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: end.Format(time.RFC3339),
		},
		Reminders: &calendar.EventReminders{
			UseDefault: false,
			Overrides: []*calendar.EventReminder{
				{Method: "popup", Minutes: 10},
			},
		},
	}

	_, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		log.Printf("Unable to create event for %s: %v", contest.Name, err)
		return
	}

	fmt.Printf("Created event: %s\n", contest.Name)
}
