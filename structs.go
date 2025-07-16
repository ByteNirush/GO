// package main

// import "fmt"

// type Person struct{
// 	Name string
// 	Age int}

// func main(){
// 	person := Person{Name: "Nirush", Age: 20}
// 	fmt.Printf("this is Person %+v\n", person)

// 	employee := struct {
// 		name string
// 		age int
// 	}{
// 		name: "Nirush",
// 		age: 20,
// 	}

// 	type Address struct {
// 		city  string
// 		state string
// 	}
// 	type Contact struct {
// 		Name 	string
// 		Address Address
// 		Phone 	string
// 	}

// 	contact := Contact{
// 		Name:    "Nirush",
// 		Address: Address{
// 			city:  "Kathmandu",
// 			state: "Bagmati",
// 		},
// 		Phone:   "1234567890",
// 	}

// 	fmt.Printf("this is contact %+v\n", contact)
// 	fmt.Println("this is employee", employee)

// 	fmt.Println("name berfore modification:", person.Name)
// 	modifyPersonName(&person)
// 	fmt.Println("name after modification:", person.Name)

// 	x := 10
// 	ptr := &x
// 	fmt.Printf("value of x: %d, and address of x %p\n", x, ptr)
// 	*ptr = 20
// 	fmt.Printf("value of x after modification: %d, and address of x %p\n", x, ptr)
// }

// func modifyPersonName(person *Person){
// 	person.Name = "Nikesh"
// 	fmt.Println("this is modified person", person.Name)

// }


package beginnergo

import "fmt"

type Person struct{
	Name string
	Age int}

func main(){
	person := Person{Name: "Nirush", Age: 20}
	fmt.Printf("this is Person %+v\n", person)

	employee := struct {
		name string
		age int
	}{
		name: "Nirush",
		age: 20,
	}

	type Address struct {
		city  string
		state string
	}
	type Contact struct {
		Name 	string
		Address Address
		Phone 	string
	}

	contact := Contact{
		Name:    "Nirush",
		Address: Address{
			city:  "Kathmandu",
			state: "Bagmati",
		},
		Phone:   "1234567890",
	}

	fmt.Printf("this is contact %+v\n", contact)
	fmt.Println("this is employee", employee)

	fmt.Println("name berfore modification:", person.Name)
	person.modifyPersonName("Nikesh")
	fmt.Println("name after modification:", person.Name)
	
}

func(p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("this is modified person", p.Name)

}