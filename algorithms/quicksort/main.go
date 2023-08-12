// Implement the quicksort algorithm and benchmark it against the standard library sort function.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Generate a slice of random numbers
	slice := generateSlice(200)
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("-----------------------------------------------------------------------------------------")

	// Benchmark the standard library sort
	start := time.Now()
	sort.Ints(slice)
	elapsed := time.Since(start)
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("\n--- Sorted ---\n\n", slice)
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("Standard library sort took:", elapsed)
	fmt.Println("-----------------------------------------------------------------------------------------")

	// Benchmark the quicksort
	slice2 := generateSlice(200)
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("\n--- Unsorted --- \n\n", slice2)
	fmt.Println("-----------------------------------------------------------------------------------------")

	start = time.Now()
	slice2 = quicksort(slice2)
	elapsed = time.Since(start)
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("\n--- Sorted ---\n\n", slice2)
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("Quicksort took:", elapsed)
	fmt.Println("-----------------------------------------------------------------------------------------")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
