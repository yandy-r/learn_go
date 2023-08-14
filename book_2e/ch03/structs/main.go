package main

import "fmt"

// Write a go struct for a car and implement some of the methods a car type would use.

type Car struct {
	make  string
	model string
	year  int
}

func (c *Car) Start() {
	println("Vroom!")
}

func (c *Car) Stop() {
	println("Screech!")
}

func (c *Car) Honk() {
	println("Beep beep!")
}

// Add a string method to the car type that returns a string representation of the car.
func (c *Car) String() string {
	return fmt.Sprintf("%d %s %s", c.year, c.make, c.model)
}

func NewCar(n string, m string, y int) *Car {
	return &Car{make: n, model: m, year: y}
}

func main() {
	myCar := NewCar("Ford", "Focus", 2013)
	fmt.Println()
	fmt.Println(myCar)
}
