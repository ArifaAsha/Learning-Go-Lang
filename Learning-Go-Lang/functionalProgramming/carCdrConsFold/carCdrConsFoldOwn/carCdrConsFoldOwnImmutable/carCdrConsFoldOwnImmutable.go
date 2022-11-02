package main

import (
	"fmt"
	"strconv"
	"strings"
)

type List struct {
	Data int
}

func (list List) SetData(val int) List {
	list.Data = val
	return list
}

func (list List) GetData() int {
	return list.Data
}

func stringToSlice(A string) []int {
	strings := strings.Split(A, " ")
	ints := make([]int, len(strings))

	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
	// fmt.Println(reflect.TypeOf(ints))
	// fmt.Printf("%#v\n", ints)
}

func convertInterfaceToSlice(list interface{}) []int {

	strConv2 := fmt.Sprintf("%v", list) //making string from interface values
	// fmt.Println(strConv2)
	convertedSlice := stringToSlice(strConv2)
	return convertedSlice
	// stringToSliceCon := stringToSlice(strConv2)
	// fmt.Println(stringToSliceCon)

}

func convertIntToSlice(x int) []int {
	slice := []int{x}
	// fmt.Println(slice)
	return slice

}

// func (list List) Cons(toAdd int) List {

// 	list2 := list.SetData(toAdd)
// 	return list2
// }

func Cons(list List, toAdd interface{}) []int {
	toAddSlice := convertInterfaceToSlice(toAdd)

	// x := 9
	// typeOfToAdd := print_out_type(toAdd)
	typeOfInt := "int"
	// fmt.Println(typeOfToAdd)

	if print_out_type(toAdd) == typeOfInt {
		interfaceToIntVal := toAdd.(int)
		intToSlice := convertIntToSlice(list)
		list = convertInterfaceToSlice(list)
		list2 := append(intToSlice, list)
		return list2

		fmt.Println(interfaceToIntVal)
		fmt.Println("hehe")
	} else {
		fmt.Println("No")
	}

	// y := print_out_type(toAdd)
	// fmt.Println(y)
	// list2 := list.SetData(toAdd)
	// return list2
}

// func Cons(list List, toAdd interface{}) []int {
// 	toAddSlice := convertInterfaceToSlice(toAdd)
// 	listToSlice := convertInterfaceToSlice(list)

// 	// // x := 9
// 	// // typeOfToAdd := print_out_type(toAdd)
// 	// typeOfInt := "int"
// 	// // fmt.Println(typeOfToAdd)

// 	// if print_out_type(toAdd) == typeOfInt {
// 	// 	interfaceToIntVal := toAdd.(int)
// 	list2 := append(listToSlice, toAddSlice)
// 	return list2

// 	// 	fmt.Println(interfaceToIntVal)
// 	// 	fmt.Println("hehe")
// 	// } else {
// 	// 	fmt.Println("No")
// 	// }

// 	// y := print_out_type(toAdd)
// 	// fmt.Println(y)
// 	// list2 := list.SetData(toAdd)
// 	// return list2
// }

// func print_out_type(x interface{}) string {
// 	return reflect.TypeOf(x).String()
// }

func main() {

	// list := List{}
	// list2 := []int{2, 3, 4}
	convertIntToSlice(5)
	// printValue(list)

	// constructedList := list.Cons(list, 9)
	// convertInterfaceToSlice(list)
	// list2 := list.Cons(2)
	// list3 := list.Cons(3)

	// fmt.Println(list2)
	// fmt.Println(list3)
}
