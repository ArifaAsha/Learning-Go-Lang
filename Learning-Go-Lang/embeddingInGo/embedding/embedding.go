package main

import "fmt"

type Person struct {
	name  string
	phone string
	add   string
}

//Repeatative fields are present
// type Employee struct {
// 	name    string
// 	id      int
// 	dept    string
// 	phone   string
// 	address string
// 	salary  int
// }

type Employee struct {
	Person //Embedding Person inside Employee
	id     int
	dept   string
	salary int
}

func (E *Employee) setInfo() {
	fmt.Scan(&E.name, &E.dept, &E.salary, &E.phone, &E.add)
}

func (E Employee) getInfo() {
	fmt.Println(E.name, E.name)
	fmt.Println(E.dept, E.dept)
}

func main() {
	emp := Employee{}
	emp.setInfo()
	emp.getInfo()
}
