package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)
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
