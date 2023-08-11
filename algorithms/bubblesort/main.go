package main

import (
	"fmt"
	"math/rand"
)

// Create a random list of ordered integers
func createList() []int {
	s := make([]int, 100)
	for i := range s {
		s[i] = rand.Intn(100)
	}

	return s
}

func bubbleSort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	us := createList()
	fmt.Println("-------------------------------------")
	fmt.Println("Unsorted list: ")
	fmt.Println(us)
	fmt.Println("-------------------------------------")
	fmt.Println()
	s := bubbleSort(us)
	fmt.Println("-------------------------------------")
	fmt.Println("Sorted list: ")
	fmt.Println(s)
	fmt.Println("-------------------------------------")
}
