package main

import "fmt"

// function closure
func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	returnFunc("hola")() //this line is replaced by the line return func() { fmt.Println(x) }
	x := returnFunc("Tata")
	x()
}
