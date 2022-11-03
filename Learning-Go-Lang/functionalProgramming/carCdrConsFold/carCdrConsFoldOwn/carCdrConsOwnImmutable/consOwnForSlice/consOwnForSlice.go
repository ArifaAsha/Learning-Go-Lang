package main

import (
	"fmt"
)

type List struct {
	Data int
}

// func stringToSlice(A string) []int {
// 	strings := strings.Split(A, " ")
// 	ints := make([]int, len(strings))

// 	for i, s := range strings {
// 		ints[i], _ = strconv.Atoi(s)
// 	}
// 	return ints
// }

// func convertInterfaceToSlice(list interface{}) []int {

// 	strConv2 := fmt.Sprintf("%v", list) //making string from interface values
// 	// fmt.Println(strConv2)
// 	if len(strConv2) > 0 {
// 		convertedSlice := stringToSlice(strConv2)
// 		return convertedSlice
// 	}
// 	return nil
// }

func convertIntToSlice(x int) []int {
	slice := []int{x}
	// fmt.Println(slice)
	return slice
}

// func makeInterfaceFromSlice(a []int) interface{} {
// 	b := make([]interface{}, len(a))
// 	for i := range a {
// 		b[i] = a[i]
// 	}
// 	return b
// }

func Cons(list []int, toAdd int) []int {
	// newList := convertInterfaceToSlice(list)

	// fmt.Println("list 1 = ", newList)
	toAddSlice := convertIntToSlice(toAdd)
	// fmt.Println("list 2 = ", toAddSlice)
	// toAddSlice := convertIntToSlice(toAdd)
	constructedList := append(list, toAddSlice[:]...)
	// fmt.Println("After concat = ", constructedList)
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
