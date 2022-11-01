package main

type Stack []int

var stack Stack

// type List struct {
// 	Data interface{}
// }

// func (List List) GetData() interface{} {
// 	return List.Data
// }

// func (List *List) SetData(data interface{}) {
// 	List.Data = data
// }

func push(toAdd int) {
	append(stack, toAdd)
}

// func car(list []int) int {
// 	head := list[0]
// 	return head
// }

// func cdr(list []int) []int {
// 	tail := list[1:]
// 	return tail
// }

func IsEmpty(s Stack) bool {
	return len(s) == 0
}

func main() {

	// list.SetData([]int{1, 2, 3, 4})
	var stack Stack
	push(10)
	println(stack)

	// fmt.Println(list2)
	// fmt.Println(car(list2))
	// fmt.Println(cdr(list2))

	// fmt.Println("---------------Stack---------------")
	// s := StaPush(list2, 433)
	// println(s)
}
