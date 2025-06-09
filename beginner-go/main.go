package main

import (
	"fmt"
)

func basicFunctions() {
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

func basicDataStructures() {
	// Control Structures
	age := 30

	if age >= 18 {
		fmt.Println("You're an adult")
	} else if age > 13 {
		fmt.Println("You're a teenager")
	} else {
		fmt.Println("You're a child")
	}

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday, start of the week!")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("It's a midweek day!")
	case "Friday":
		fmt.Println("It's Friday, almost the weekend!")
	default:
		fmt.Println("It's the weekend!")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("This is i", i)
	}

	counter := 0

	for counter < 3 {
		fmt.Println("Counter is", counter)
		counter++
	}

	iterations := 0

	// Infinite loop with a break condition
	for {
		if iterations >= 3 {
			break
		}
		fmt.Println("Iteration number", iterations)
		iterations++
	}

	// Arrays and Slices
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array of numbers: %v\n", numbers)
	fmt.Println("This is the last element of the array:", numbers[len(numbers)-1])

	// numbersAtInt := [...]int{1, 2, 3, 4, 5}

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("Matrix: %v\n", matrix)

	// allNumbers := numbers[:]
	// firstThreeNumbers := numbers[0:3]

	// allNumbers = append(allNumbers, 60, 70, 80)

	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("These are my fruits: %v\n", fruits)

	fruits = append(fruits, "Kiwi")
	fmt.Printf("These are my fruits: %v\n", fruits)

	fruits = append(fruits, "Orange", "Grapes")
	fmt.Printf("These are my fruits: %v\n", fruits)

	moreFruits := []string{"Mango", "Pineapple"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("These are my fruits: %v\n", fruits)

	for i, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	capitalCities := map[string]string{
		"Germany": "Berlin",
		"France":  "Paris",
		"Spain":   "Madrid",
	}

	fmt.Printf("Capital of Germany: %s\n", capitalCities["Germany"])
	capital, exists := capitalCities["Italy"]

	if exists {
		fmt.Printf("Capital of Italy: %s\n", capital)
	} else {
		fmt.Println("Capital of Italy not found")
	}

	delete(capitalCities, "Spain")
	fmt.Printf("Capital cities after deletion: %v\n", capitalCities)
}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

type Person struct {
	Name string
	age  int
}

func main() {
	person := Person{Name: "Artem", age: 27}
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Person's name: %s, age: %d\n", person.Name, person.age)

	employee := struct {
		name string
		id   string
	}{
		name: "John Doe",
		id:   "E12345",
	}

	fmt.Printf("Employee: %+v\n", employee)

	type Address struct {
		Street string
		City   string
	}
	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "Alice",
		Address: Address{
			Street: "123 Main St",
			City:   "Wonderland",
		},
		Phone: "123-456-7890",
	}
	fmt.Printf("Contact: %+v\n", contact)

	fmt.Println("Old person's name:", person.Name)
	person.modifyPersonName("Alex")
	fmt.Println("New person's name:", person.Name)

	x := 20
	ptr := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)

	*ptr = 30 // Modifying the value at the address pointed to by ptr
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)
}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("Inside modifyPersonName function, person's name:", p.Name)
}
