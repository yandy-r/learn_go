package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

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

func getTarget() (int, error) {
	var s string
	if len(os.Args) > 1 {
		s = os.Args[1]
	} else {
		err := fmt.Errorf("No target value entered")
		return -1, err
	}

	// convert s to int and return it
	t, e := strconv.Atoi(s)
	if e != nil {
		err := fmt.Errorf("Error: %v, entered value: %v", t, reflect.TypeOf(t))
		return -1, err
	}

	return t, nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	t, err := getTarget()
	i := binary_search(arr, t)
	fmt.Println()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("-------------------------------------")
	if i != -1 {
		fmt.Printf("The target: %d is at index: %d\n", t, i)
	} else {
		fmt.Printf("The target: %d is not in the array\n", t)
	}
	fmt.Println("-------------------------------------")
	fmt.Println()
}
