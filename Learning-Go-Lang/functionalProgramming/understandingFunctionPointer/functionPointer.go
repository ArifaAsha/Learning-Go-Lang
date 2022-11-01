package main

import "fmt"

func test(x int) {
	fmt.Println("Hello!", x)
}

func main() {
	x := test //reference the function; assign function to a variable
	x(4)      //works same as statement => test()
}
