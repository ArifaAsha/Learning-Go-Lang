package main

import "fmt"

func callPrintNumber(f func(string), number string) { //function as a parameter
	f(number)
}

func printNumber(number string) {
	fmt.Println("The number is: ", number)
}

// Closure ->
func getClicker() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	callPrintNumber(printNumber, "12222") //Function as a parameter
	click := getClicker()                 //returning function
	fmt.Println(click())
	fmt.Println(click())
}
