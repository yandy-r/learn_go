package main

import (
	"fmt"
	"os"
)

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

// Create a function to get command line arguments and print them
func printArgs() {

	if len(os.Args) < 2 {
		e := fmt.Errorf("At least one argument is required")

		fmt.Println(e)
	}

	fmt.Println("Arguments:", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1 // this will not compile
	// y = "bye" // this will not compile

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println()
	printArgs()
}
