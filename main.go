package main

import (
	"fmt"
)

func main() {
	// Variable 
	var name string = "Nirush"
	fmt.Printf("This is my name %s\n", name)

	age:= 40
	fmt.Printf("This is my age %d\n", age)

	// var city string
	// city = "KTM"
	// fmt.Printf("This is my city %s\n", city)	

	var country, continent string = "Nepal", "Asia"
	fmt.Printf("This is my country %s and This is my continent %s\n", country, continent)

	var (
		isEmployed bool = true
		salary int = 5000
		position string = "Software Engineer"
	)
	fmt.Printf("isEmployed: %t, salary: %d, and position: %s\n", isEmployed, salary, position)

	// zero values
	var defaultInt int
	var defaultFloat float64 
	var defaultString string
	var defaultBool bool

	fmt.Printf("int: %d, float: %f, string: '%s', bool: %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	// Constants
	const (
		monday = 2
		tuesday = 3
		wednesday = 4
	)

	fmt.Printf("Monday: %d, Tuesday: %d, Wednesday: %d\n", monday, tuesday, wednesday)

	const typedAge int = 19
	const untypedAge = 19
	fmt.Println(typedAge == untypedAge)

	const (
		Jan = iota + 1 
		Feb
		Mar
		Apr
	)

	fmt.Printf("Jan - %d, Feb - %d, Mar - %d, Apr - %d\n", Jan, Feb, Mar, Apr)

	type Role int
	const (
	Admin Role = iota
	Editor
	Viewer
	)
	fmt.Printf("Admin: %d, Editor: %d, Viewer: %d\n", Admin, Editor, Viewer)


	// Functions
	resutlt := add(3, 4)
	fmt.Println("This is the result", resutlt)

	sum, product := calculateSumAndProduct(5, 6)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)


	// conditional statements
	// age := 15

	// if age >= 18{
	// 	fmt.Println("You are an adult.")
	// } else if age >= 13 {
	// 	fmt.Println("You are a teenager.")
	// } else {
	// 	fmt.Println("You are a child.")
	// }


	// switch statement

	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("It's a midweek day.")
	case "Friday":
		fmt.Println("It's almost the weekend.")
	default:
		fmt.Println("It's the weekend!")
	}


	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 0
	for counter < 3 {
		fmt.Println("Counter:", counter)
		counter++
	}

	iteration := 0
	for  {
		if iteration >= 5 {
			break
		}	
		//some condition is meet
		iteration++
	}


	// Arrays and Slices
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("This is an array: %v\n", numbers)

	fmt.Println("This is the last value", numbers[len(numbers)-3])


	// numbersAtInit := [...]int{1, 2, 3, 4, 5}

	matrix := [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Printf("This is a 2D array: %v\n", matrix)

	// allNumbers := numbers[:]
	// firstThree := numbers[0:3]


	// Slices
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("this is a fruits %v\n", fruits)

	fruits = append(fruits, "Date")
	fmt.Printf("After appending, fruits: %v\n", fruits)

	fruits = append(fruits, "Mango")
	fmt.Printf("After appending again, fruits: %v\n", fruits)

	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Maps
	capitalCities := map[string]string{
		"Nepal": "Kathmandu",
		"USA": "Washington, D.C.",
		"UK": "London",
	}
	fmt.Println(capitalCities["Nepal"])
	fmt.Println(capitalCities["USA"])

	
	capital, exists := capitalCities["Poland"]
	if exists {
		fmt.Println("This is the capital", capital)
	} else {
		fmt.Println("This capital is not found.")
	}

	// delete
	delete(capitalCities, "UK")
	fmt.Println("After deletion, capitalCities:", capitalCities)

}

// functions
func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}