package main

import "fmt"

func main() { 
	// fmt.Println("Hello, Codeforces Toolkit!")
	fmt.Println("Fetching contests...")
	data := getContests()
	fmt.Println(data)
}