package main

import (
	"fmt"
	"log"
	"time"
)

func main() { 
	// fmt.Println("Hello, Codeforces Toolkit!")
	fmt.Println("Fetching contests...")
	contests, err := getContests()

	if err != nil {
		log.Fatal(err)
	}

	for _, c := range contests {
		t := time.Unix(c.StartTimeSeconds, 0)
		// fmt.Printf("%s (ID: %d)\n", c.Name, c.ID)
		fmt.Printf("- %s\n When: %s\n\n", c.Name, t.Format(time.RFC1123))
	}

	fmt.Println("\nInitializing Google Calendar...")
	srv := getCalendarService()
	fmt.Println("Success! Authenticated with Google Calendar.")

	fmt.Printf("Service: %v\n", srv)
}