package main

import (
	"flag"
	"fmt"
	"strings"
)

var f string

func init() {
	flag.StringVar(&f, "function", "", "The function to run.")
	flag.StringVar(&f, "f", "", "The function to run (shorthand).")
}

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
	flag.Parse()
	switch strings.ToLower(f) {
	case "commaok", "comma_ok":
		commaOk()
	}
}
