package main

import (
	"fmt"
	"os"
	"strings"
)

/*
	 The comma, ok idiom is used to test for the existence of a key in a map.
	 Output:
	    5 true

		0 true

		0 false
*/
func commaOk() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println()
	v, ok := m["hello"]
	fmt.Println(v, ok)
	fmt.Println()

	v, ok = m["world"]
	fmt.Println(v, ok)
	fmt.Println()

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
	fmt.Println()

}

func main() {
	args := os.Args[1:]

	switch arg := args[0]; strings.ToLower(arg) {
	case "commaok", "comma_ok":
		commaOk()
	}
}
