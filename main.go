package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	mode := flag.String("mode", "dev", "Mode to run: dev, web, or worker")
	flag.Parse()

	initDB()
	initOAuth()

	switch *mode {
	case "worker":
		fmt.Println("Running in WORKER mode...")
		runSyncAllUsers()
	case "web":
		fmt.Println("Running in WEB mode...")
		startServer() // Blocks
	case "dev":
		fmt.Println("Running in DEV mode...")
		go startServer()

		// Run sync immediately
		runSyncAllUsers()

		// Loop
		ticker := time.NewTicker(24 * time.Hour)
		for range ticker.C {
			runSyncAllUsers()
		}
	}
}

func runSyncAllUsers() {
	fmt.Println("Starting sync job...")

	// 1. Fetch contests ONCE for everyone
	contests, err := getContests()
	if err != nil {
		log.Printf("Error fetching contests: %v", err)
		return
	}
	fmt.Printf("Found %d upcoming contests.\n", len(contests))

	// 2. Get all users
	users, err := getAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return
	}

	// 3. Sync for each user
	for _, u := range users {
		fmt.Printf("Syncing for user: %s\n", u.Email)

		// Reconstruct the token
		token := &oauth2.Token{
			AccessToken:  u.AccessToken,
			RefreshToken: u.RefreshToken,
			Expiry:       u.TokenExpiry,
			TokenType:    "Bearer",
		}

		// Create a client for THIS user
		// We use googleConfig.TokenSource to handle auto-refreshing!
		tokenSource := googleConfig.TokenSource(context.Background(), token)
		client := oauth2.NewClient(context.Background(), tokenSource)

		srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Printf("  Failed to create calendar service: %v", err)
			continue
		}

		// Add events
		for _, c := range contests {
			// TODO: We need a better way to check if *this specific user* has the event.
			// For now, we'll just try to add it. Google Calendar handles duplicates gracefully-ish
			// (it won't error, but might create doubles if we aren't careful).
			// Ideally, we'd check srv.Events.List().

			addContestToCalendar(srv, c)
		}
	}
}
