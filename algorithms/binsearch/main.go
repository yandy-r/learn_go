package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
NOTE: Struct Node is not necessary for binary search, but I wanted to use it to
practice creating structs and methods. The struct contains three fields:
high, low, and mid. The high and low fields are the high and low indexes of
the array that the node represents. The mid field is the average of the high
and low fields. The mid field is the index that is checked against the target value.
*/

// Node is a struct that represents a node in a binary search tree.
type Node struct {
	high int
	low  int
	mid  int
}

// NewNode creates a new node with the given high and low values and sets the
// mid value to the average of the high and low values.
func NewNode(high, low int) *Node {
	return &Node{high: high, low: low, mid: (high + low) / 2}
}

// String returns a string representation of the node.
func (n *Node) String() string {
	return fmt.Sprintf("high: %d, low: %d, mid: %d", n.high, n.low, n.mid)
}

func binary_search(arr []int, target int) int {
	item := NewNode(len(arr)-1, 0)

	for item.low <= item.high {
		if arr[item.mid] == target {
			return item.mid
		} else if arr[item.mid] > target {
			item.high = item.mid - 1
		} else if arr[item.mid] < target {
			item.low = item.mid + 1
		}
		item.mid = (item.low + item.high) / 2
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
		err := fmt.Errorf("Cannot convert \"%s\" to int", s)
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
