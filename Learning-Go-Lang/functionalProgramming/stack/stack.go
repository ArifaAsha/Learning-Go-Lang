package main

import "fmt"

type Any interface{}
type FuncType func(Any, Any) Any
type ConsType func(FuncType) Any

func Cons(a, b Any) ConsType { //construct -> pairs
	return func(fn FuncType) Any {
		return fn(a, b)
	}
}
func Car(cons ConsType) Any { //First element of the list = heaD
	head := func(a, b Any) Any { return a }
	return cons(head)
}

func Cdr(cons ConsType) Any { //Rest elements of the list = Tail
	tail := func(a, b Any) Any { return b }
	return cons(tail)
}

// func walkDownList(list ConsType) {
// 	var printer func(a, b Any) Any
// 	printer = func(a, b Any) Any {
// 		fmt.Println(a)
// 		if b == nil {
// 			fmt.Println('m')
// 			return nil
// 		}
// 		if bVal, castOk := b.(ConsType); castOk {
// 			bVal(printer)
// 		}
// 		return nil
// 	}
// }

// func check_empty(list ConsType)
//     return nil

// Adding items into the stack
func push(list ConsType, item int) {
	b := make([]Any, 0, 5)
	appendedList := append(b, Cons(list, item))

	for _, num := range appendedList {
		fmt.Println(num)
	}
	// println(appendedList)

	// println("pushed item: ", item)
}

// func stack(fn FuncType, init Any, list ConsType) Any {
// 	if list == nil {
// 		return init
// 	}
// 	head, tail := Car(list), Cdr(list)
// 	rest, _ := tail.(ConsType)

// 	return stack(fn, (fn(init, head)), rest)
// }

// func stackInGo(list ConsType){
// 	l := len(list)
// 	if list == nil{
// 		fmt.Println("Null List")
// 		return list
// 	}
// }

// func walkDownList(list ConsType) {
// 	var printer func(a, b Any) Any
// 	printer = func(a, b Any) Any {
// 		fmt.Println(a)
// 		if b == nil {
// 			fmt.Println('m')
// 			return nil
// 		}
// 		if bVal, castOk := b.(ConsType); castOk {
// 			bVal(printer)
// 		}
// 		return nil
// 	}
// 	list(printer)
// }

func main() {
	a := Cons(12, 13)

	fmt.Println("Hello, playground")
	// walkDownList(a)
	push(a, 1)
	// walkDownList(a)

	// fmt.Println(a(add2))
	// list := Cons(12, Cons(14, Cons(16, nil)))
	// fmt.Println(list, '!')
	// // walkDownList(list)
	// fmt.Println("..", foldl(add2, 0, list))
	// fmt.Println("!!", foldr(add2, 0, list))
}
