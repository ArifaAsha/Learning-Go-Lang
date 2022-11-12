package main

import (
	"fmt"
)

type List struct {
	Data int
}

func convertIntToSlice(x int) []int {
	slice := []int{x}
	return slice
}

func Cons(list []int, toAdd int) []int {
	toAddSlice := convertIntToSlice(toAdd)
	constructedList := append(list, toAddSlice[:]...)
	return constructedList
}

func main() {

	var list []int
	list2 := Cons(list, 7)
	fmt.Println(list2)

	list3 := Cons(list2, 9)
	fmt.Println(list3)

	list4 := Cons(list3, 10)
	fmt.Println(list4)

}

/*
Output:

[7]
[7 9]
[7 9 10]
*/
