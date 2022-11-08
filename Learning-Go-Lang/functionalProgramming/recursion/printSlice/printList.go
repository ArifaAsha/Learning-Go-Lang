package main

func printList(slice []int) {
	if len(slice) == 0 {
		return
	} else {
		print(slice[0])
		printList(slice[1:])
	}
}

func main() {
	l := []int{1, 2, 3}
	printList(l)
}
