package main

import "fmt"

func fitness(candidate string) int {
	count := 0
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == '1' { // count the one-bits (a single character)
			count++
		}
	}
	return count
}

func main() {
	a, b := "1001", "0110"

	// Crossover: the first two bits of a, then the rest of b.
	child := a[:2] + b[2:4]

	// Mutation: change the second bit of the child to 1.
	mutated := child[:1] + "1" + child[2:]

	fmt.Println("Parent A:", "fitness:", fitness(a))
	fmt.Println("Parent B:", "fitness:", fitness(b))
	fmt.Println("Child:", child, "fitness:", fitness(child))
	fmt.Println("Mutated:", mutated, "fitness:", fitness(mutated))
}
