package main

import (
	"fmt"
	"math/rand"
)

// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
// Loop over the slice created. For each value in the slice, apply the following rules:
// - if the value is divisiable by 2, print Two!
// - if the value is divisable by 3, print Three!
// - if the valu is divisible by 2 and 3, print Six!
// - Otherwise print Never Mind
func main() {
	s := make([]int, 100)
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(100))
	}

	fmt.Println()
	for _, v := range s {
		if v%3 == 0 && v%2 == 0 {
			fmt.Println("Six!")
		} else if v%3 == 0 {
			fmt.Println("Three!")
		} else if v%2 == 0 {
			fmt.Println("Two!")
		} else {
			fmt.Println("Never Mind")
		}
	}
}
