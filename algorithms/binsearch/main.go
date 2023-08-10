package main

import "fmt"

func binary_search(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 3
	index := binary_search(arr, target)
	fmt.Println()
	fmt.Printf("The target: %d is at index: %d\n", target, index)
	fmt.Println()
}
