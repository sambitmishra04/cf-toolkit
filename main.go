package main

import (
	"fmt"
	"log"
	"time"
)

func main() { 

	initDB()
	// fmt.Println("Hello, Codeforces Toolkit!")
	fmt.Println("Fetching contests...")
	contests, err := getContests()

	if err != nil {
		log.Fatal(err)
	}

	
	fmt.Println("\nInitializing Google Calendar...")
	srv := getCalendarService()
	fmt.Println("Success! Authenticated with Google Calendar.")
	
	// fmt.Printf("Service: %v\n", srv)

	for _, c := range contests {

		if eventExists(c.ID) {
			fmt.Printf("Skipping %s (already added)\n", c.Name)
			continue
		}
		t := time.Unix(c.StartTimeSeconds, 0)
		// fmt.Printf("%s (ID: %d)\n", c.Name, c.ID)
		fmt.Printf("- Adding %s\n  When: %s\n", c.Name, t.Format(time.RFC1123))

		// fmt.Printf("- %s\n When: %s\n\n", c.Name, t.Format(time.RFC1123))

		addContestToCalendar(srv, c)
		// fmt.Println("Test mode: Added 1 event and stopping.")
		// break

		saveEvent(c.ID)

	}
}