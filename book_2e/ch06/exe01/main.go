package main

import "fmt"

/*
   Create a struct named Person with three fields:
   - FirstName string
   - lastName string
   - Age int
   Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person.
   Write another function calle MakePersonPointer that takes in firstName, lastName, and age and returns a pointer to a Person.
   run "go build -gcflags="-m"" to see the escape analysis.
*/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	p := MakePerson("John", "Doe", 30)
	pt := MakePersonPointer("John", "Doe", 30)

	fmt.Println()
	fmt.Println(p)
	fmt.Println(pt)
	fmt.Println()
}
