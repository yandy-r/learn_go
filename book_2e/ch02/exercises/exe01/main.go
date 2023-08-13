package main

import "fmt"

/* Write a program where you declare an integer variable called i with the value 20.
Assign i to a floating point variable named f. Print out i and f.
*/

func main() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println()
	fmt.Printf("i = %d, f = %f\n", i, f)
}
