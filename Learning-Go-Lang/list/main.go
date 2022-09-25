package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront(10)
	l.PushFront(20)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}

	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:2]
	y := arr[1:3]
	fmt.Print(x)
	fmt.Println(y)

}
