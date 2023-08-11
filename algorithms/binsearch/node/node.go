package node

import "fmt"

/*
NOTE: Struct Node is not necessary for binary search, but I wanted to use it to
practice creating structs and methods. The struct contains three fields:
high, low, and mid. The high and low fields are the high and low indexes of
the array that the node represents. The mid field is the average of the high
and low fields. The mid field is the index that is checked against the target value.
*/

// node is a struct that represents a node in a binary search tree.
type node struct {
	high int
	low  int
	mid  int
}

// NewNode creates a new node with the given high and low values and sets the
// mid value to the average of the high and low values.
func NewNode(high, low int) *node {
	return &node{high: high, low: low, mid: (high + low) / 2}
}

// String returns a string representation of the node.
func (n *node) String() string {
	return fmt.Sprintf("high: %d, low: %d, mid: %d", n.high, n.low, n.mid)
}

func (n *node) GetHigh() int {
	return n.high
}

func (n *node) GetLow() int {
	return n.low
}

func (n *node) GetMid() int {
	return n.mid
}

func (n *node) BinSearch(arr []int, t int) int {
	for n.low <= n.high {
		if arr[n.mid] == t {
			return n.mid
		} else if arr[n.mid] > t {
			n.high = n.mid - 1
		} else if arr[n.mid] < t {
			n.low = n.mid + 1
		}
		n.mid = (n.low + n.high) / 2
	}
	return -1
}
