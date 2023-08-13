/* Write a program with three variables, one named b of type byte,
   one named smallI of type int32, and one named bigI of type uint64.
   Assign each variable the maximum legal value for its type,
   then add 1 to each variable. Print out their values.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	fmt.Println()
	fmt.Printf("b = %v, smallI = %v, bigI = %v\n", b, smallI, bigI)

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println()
	fmt.Printf("b = %v, smallI = %v, bigI = %v\n", b, smallI, bigI)
	fmt.Println()
}
