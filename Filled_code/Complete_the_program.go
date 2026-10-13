package main

import "fmt"

var taxRate = 0.1

func main() {
	price := 20000.0
	qty := 3
	subtotal := price * float64(qty)
	total := subtotal + subtotal*taxRate
	isBig := total > 50000
	fmt.Println(total, isBig)
}