package main

import "fmt"

func main() {
	first, second := 1, 1

	// Each new number is the sum of the previous two.
	third := first + second
	fourth := second + third
	fifth := third + fourth

	fmt.Println(third, fourth, fifth)
}