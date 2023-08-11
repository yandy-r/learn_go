package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"github.com/kidinstinct/learn_go/algorithms/binsearch/node"
)

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
		err := fmt.Errorf("Cannot convert \"%s\" to int", s)
		return -1, err
	}

	return t, nil
}

// Create a random list of ordered integers
func orderedList() []int {
	s := make([]int, 100)
	for i := range s {
		s[i] = rand.Intn(100)
	}
	sort.Ints(s)

	return s
}

func main() {
	arr := orderedList()
	t, err := getTarget()

	n := node.NewNode(len(arr)-1, 0)
	i := n.BinSearch(arr, t)

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
