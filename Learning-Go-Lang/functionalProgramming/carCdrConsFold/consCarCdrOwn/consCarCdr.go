package main

import "fmt"

type cons struct {
	a, b interface{}
}

func car(fun func) cons:
    func cons(a,b):
        return a
    return fun(cons)

func main() {
	p1 := cons{1, 42}
	fmt.Println("p1=", p1)

}
