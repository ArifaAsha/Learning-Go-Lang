package main

import "fmt"

func callPrintNumber(f func(string), number string) {
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
	callPrintNumber(printNumber, "12222")
	click := getClicker()
	fmt.Println(click())
	fmt.Println(click())
}
