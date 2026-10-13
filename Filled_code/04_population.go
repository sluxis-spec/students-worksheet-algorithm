package main

import "fmt"

func main() {
	start, births, arrivals, deaths, departures := 1000, 20, 5, 10, 4

	// final = initial + births + arrivals - deaths - departures
	end := start + births + arrivals - deaths - departures

	fmt.Println("Population:", end)
}