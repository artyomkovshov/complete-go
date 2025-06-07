package main

import (
	"fmt"
)

func main() {
	// Variables

	var name string = "Artem"
	fmt.Printf("This is my name: %s\n", name)

	age := 27
	fmt.Printf("This is my age: %d\n", age)

	var city string
	city = "Berlin"
	fmt.Printf("This is my city: %s\n", city)

	var country, continent string = "Germany", "Europe"
	fmt.Printf("This is my country: %s, and this is my continent: %s\n", country, continent)

	var (
		isEmployed bool   = true
		salary     int    = 100000
		position   string = "Developer"
	)

	fmt.Printf("Is employed: %t, this is my salary: %d, this is my position: %s\n", isEmployed, salary, position)

	// Zero (default) values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("Default int: %d, Default float: %f, Default string: '%s', Default bool: %t\n",
		defaultInt, defaultFloat, defaultString, defaultBool)

	// Constants
	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Printf("Pi: %f, Monday: %d, Tuesday: %d, Wednesday: %d\n", pi, Monday, Tuesday, Wednesday)

	const typedAge int = 30
	const untypedAge = 30
	fmt.Printf("Typed age: %d, Untyped age: %d\n", typedAge, untypedAge)

	const (
		Jan int = iota + 1 // 1
		Feb                // 2
		Mar                // 3
		Apr                // 4
	)

	fmt.Printf("Months: Jan: %d, Feb: %d, Mar: %d, Apr: %d\n", Jan, Feb, Mar, Apr)

	result := add(3, 4)
	fmt.Printf("Result of add function: %d\n", result)

	sum, product := calculateSumAndProduct(10, 11)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
