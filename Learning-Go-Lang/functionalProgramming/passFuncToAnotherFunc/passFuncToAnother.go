package main

import "fmt"

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func main() {
	test :=
		func(x int) int {
			return x * -4
		}

	test3 :=
		func(x int) int {
			return x * 7
		}
	test2(test)
	test2(test3)
}
