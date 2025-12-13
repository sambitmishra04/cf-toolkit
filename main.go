package main

import (
	"fmt"
	"log"
)

func main() { 
	// fmt.Println("Hello, Codeforces Toolkit!")
	fmt.Println("Fetching contests...")
	contests, err := getContests()

	if err != nil {
		log.Fatal(err)
	}

	for _, c := range contests {
		fmt.Printf("%s (ID: %d)\n", c.Name, c.ID)
	}
}