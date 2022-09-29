package main

import "fmt"

// defer adds statements in stack --> LIFO
func main() {
	defer fmt.Println(1)
	defer fmt.Println(2) //defer -> skipped. Before returning the value, defer statement is executed
	defer fmt.Println(3)
	fmt.Println("Hola") // 1. Print
	myDefer()           // 2. non defer --> will be called first
}

// 3 2 1
// 0, 1, 2, 3, 4
// Hola, 43210, 3, 2, 1

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i) //
	}
}
