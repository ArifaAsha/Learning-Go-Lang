package main

import "fmt"

type Person struct {
	name     string
	phone    string
	location string
}

//function
// func print_data(){
// ...
// }

//method -> function with receiver argument

// value receiver
func (p Person) print_persons() { //p ==> per
	p.name = "Changed"
	fmt.Println("Name: ", p.name)
	fmt.Println("Phone: ", p.phone)
	fmt.Println("Location: ", p.location)
}

func main() {
	per := Person{"A", "0177929920", "Dhaka"}
	per.print_persons() //passed to the value receiver
	fmt.Println("Name in the caller function: ", per.name)
}
