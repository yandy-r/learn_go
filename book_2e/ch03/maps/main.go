package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func initFlags(args []string) (*flag.FlagSet, error) {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)

	fs.String("function", "", "The function to run.")
	fs.String("f", "", "The function to run (shorthand).")
	fs.Bool("v", false, "Verbose")

	err := fs.Parse(args)
	return fs, err
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
	f, e := initFlags(os.Args[1:])
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	if os.Args[1] != f.Name() {
		e = errors.New("Expected run flag set")
		fmt.Println()
		fmt.Println(e)
		fmt.Println()
		os.Exit(1)
	}

	subCmd := f.Arg(2)
	fmt.Println()
	fmt.Println(subCmd)
	switch strings.ToLower(subCmd) {
	case "commaok", "comma_ok":
		commaOk()
	}
}
