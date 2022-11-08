package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * (n - 1)

}

func main() {
	n := 5
	fmt.Println(factorial(n))
}
