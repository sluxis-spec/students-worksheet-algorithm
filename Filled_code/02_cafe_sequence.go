package main

import "fmt"

func main() {
	unitPrice := 18000
	quantity := 2
	discount := 5000

	// Order matters: compute the subtotal BEFORE the discount.
	subtotal := unitPrice * quantity
	total := subtotal - discount // subtotal after the discount

	fmt.Println("Subtotal:", subtotal)
	fmt.Println("Total:", total)
}