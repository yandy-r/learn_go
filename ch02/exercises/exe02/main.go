/* Write a program where you declare a constant called value that can be assigned
to both an integer and a floating point variable. Assign it to an integer called i
and floating point variable called f. Print out i and f.
*/

package main

import "fmt"

const value = 40

func main() {
	var i = value
	var f float64 = value

	fmt.Println()
	fmt.Printf("i = %d, f = %f\n", i, f)
}
