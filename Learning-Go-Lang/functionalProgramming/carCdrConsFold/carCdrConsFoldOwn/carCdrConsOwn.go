package main

import "fmt"

type List struct {
	Data interface{}
}

func (List List) GetData() interface{} {
	return List.Data
}

func (List *List) SetData(data interface{}) {
	List.Data = data
}

func cons(list *List, toAdd int) []int {
	var data []int = list.GetData().([]int)
	list2 := append(data, toAdd)
	return list2
}

// func cons(list *List, toAdd int) []int {
// 	var data []int = list.GetData().([]int)
// 	list2 := append(data, toAdd)
// 	fmt.Println(list2)
// 	return list2
// }

// // func cons(list *List, toAdd int) []int {
// func cons(list *List, toAdd int) {
// 	// list2 := &List{}
// 	var data []int = list.GetData().([]int)

// 	list2.SetData(toAdd)
// 	// fmt.Println(list2)
// 	// var listOfToAdd []int = list2.GetData().([]int)
// 	fmt.Println(list2)
// 	return append(list2, toAdd)
// 	// return list2
// }

// func cons(list Any, toAddInList int) {
// 	f := &Any{}
// 	f.SetData([]int{list})
// 	fmt.Println(f)

// 	// list2 := append(s, toAddInList)
// 	// fmt.Println(list2)
// 	// return list2

// 	// list* = []int{}
// 	// func() {
// 	// 	list2 = append(list, toAddInList)
// 	// }
// 	// return list2

// }

// func InterfaceSlice(slice interface{}) []interface{} {
// 	s := reflect.ValueOf(slice)
// 	if s.Kind() != reflect.Slice {
// 		panic("InterfaceSlice() given a non-slice type")
// 	}

// 	// Keep the distinction between nil and empty slice input
// 	if s.IsNil() {
// 		return nil
// 	}

// 	ret := make([]interface{}, s.Len())

// 	for i := 0; i < s.Len(); i++ {
// 		ret[i] = s.Index(i).Interface()
// 	}

// 	return ret
// }

// func cons(list interface{}, toAddInList int) []int {
// 	s := InterfaceSlice(list)
// 	// fmt.Println("%T", s)

// 	list2 := append(s, toAddInList)
// 	// fmt.Println(list2)
// 	// return list2

// 	// list* = []int{}
// 	// func() {
// 	// 	list2 = append(list, toAddInList)
// 	// }
// 	return list2

// }

func car(list []int) int {
	head := list[0]
	return head
}

func cdr(list []int) []int {
	tail := list[1:]
	return tail
}

func main() {
	list := &List{}
	list.SetData([]int{1, 2, 3, 4})

	// list2 := cons(list, 9)
	// cons((cons(list, 9), 10))

	list2 := cons(list, 10)
	// list2 = list.GetData().([]int)

	// var data []int = list.GetData().([]int)
	// var data []int = list.GetData().([]int)
	// fmt.Println(data)
	fmt.Println(list2)

}
