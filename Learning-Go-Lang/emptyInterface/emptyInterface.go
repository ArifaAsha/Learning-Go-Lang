package main

import "fmt"

type emptyInterface interface {
}

func main() {
	var empty interface{}
	empty = 5
	fmt.Println("Empty with 5: ", empty)

	empty = "hehe"
	fmt.Println("Empty with hehe: ")

	empty = []int{1, 2, 3}
	fmt.Print("Slice: ", empty)
	// fmt.Println(len(empty)) //interface cant be used directly

}
