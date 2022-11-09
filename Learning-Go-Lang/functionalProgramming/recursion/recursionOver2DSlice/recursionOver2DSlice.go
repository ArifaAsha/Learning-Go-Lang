package main

func printList(slice []int) {
	if len(slice) == 0 {
		return
	} else {
		print(slice[0])
		printList(slice[1:])

		// fmt.Println(slice[1:])
		// fmt.Println("==============")
	}
}

func print2DList(slice [][]int) {
	if len(slice) == 0 {
		return
	} else {
		printList(slice[0])
		print()
		print2DList(slice[1:])
	}
}

func main() {
	a := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
	}
	print2DList(a)
}
