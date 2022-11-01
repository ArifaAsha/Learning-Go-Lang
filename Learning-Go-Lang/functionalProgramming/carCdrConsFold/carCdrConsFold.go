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

/*
foldl => Fold from left
foldl :: (a -> b -> a) -> a -> [b] -> a
foldl f z []     = z
foldl f z (x:xs) = foldl f (f z x) xs
*/
func foldl(fn FuncType, init Any, list ConsType) Any {
	if list == nil {
		return init
	}
	head, tail := Car(list), Cdr(list)
	rest, _ := tail.(ConsType)

	return foldl(fn, (fn(init, head)), rest)
}

/*
foldl => Fold from right
foldr :: (a -> b -> b) -> b -> [a] -> b
foldr f z []     = z
foldr f z (x:xs) = f x (foldr f z xs)
*/
func foldr(fn FuncType, fini Any, list ConsType) Any {
	if list == nil {
		return fini
	}

	head, tail := Car(list), Cdr(list)
	rest, _ := tail.(ConsType)
	return fn(head, foldr(fn, fini, rest))
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
// 	list(printer)
// }

func main() {
	fmt.Println("Hello, playground")
	a := Cons(12, 13)
	add2 := func(x, y Any) Any {
		fmt.Println("...", x, y)
		return x.(int) + y.(int)
	}
	fmt.Println(a(add2))
	list := Cons(12, Cons(14, Cons(16, nil)))
	fmt.Println(list, '!')
	walkDownList(list)
	fmt.Println("..", foldl(add2, 0, list))
	fmt.Println("!!", foldr(add2, 0, list))
}
