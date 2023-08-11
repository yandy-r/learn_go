package main

import (
	"fmt"
	"os"
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

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
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
