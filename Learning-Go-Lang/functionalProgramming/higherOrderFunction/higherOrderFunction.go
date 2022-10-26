package main

import "fmt"

func callPrintNumber(f func(string), number string) {
	f(number)
}

func printNumber(number string) {
	fmt.Println("The number is: ", number)
}

func main() {
	callPrintNumber(printNumber, "12222")
}
