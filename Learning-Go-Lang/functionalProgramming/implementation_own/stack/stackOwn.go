package main

import "fmt"

const MAX = 100

type stack struct {
	top         int
	actualStack []int
}

func (s stack) IsEmpty() bool {
	return s.top < 0
}

// LIFO -> add as the last element of slice -> push
// Iterative approach

// func (s stack) Push(vals ...int) (stack, []int) {
// 	for _, val := range vals {
// 		s.top += 1
// 		s.actualStack = append(s.actualStack, val)
// 	}
// 	return s, s.actualStack
// }

func noOfInToBePushed(vals ...int) int {
	emptySlice := []int{}
	s := append(emptySlice, vals...)
	n := len(s)
	return n
}

// LIFO -> add as the last element of slice -> push
func (s stack) Push(vals ...int) (stack, []int) {
	noOfArgs := noOfInToBePushed(vals...)
	s.top = s.top + noOfArgs
	s.actualStack = append(s.actualStack, vals...)
	return s, s.actualStack
}

// LIFO -> last element of slice -> popped
func (s stack) Pop() int {
	if s.top < 0 {
		return 0
	}
	poppedVal := s.actualStack[s.top]
	s.top--
	return poppedVal
}

func (s stack) Top() int {
	if !s.IsEmpty() {
		val := s.actualStack[s.top]
		return val
	}
	return 0
}

func main() {
	s := stack{top: -1}
	stk, actaulStk := s.Push(1, 3, 4, 6)
	fmt.Println(actaulStk)

	x := 9
	stkNewPush, actualStkAfterPush := stk.Push(x)
	fmt.Println("New stack object after push: ", stkNewPush)
	fmt.Println("Stack after pushing 9: ", actualStkAfterPush)

	stk.Pop()
	fmt.Println(stk)
	fmt.Println("After pop: ", stk.actualStack)

	topVal := stk.Top()
	fmt.Println("Top value: ", topVal)
}
