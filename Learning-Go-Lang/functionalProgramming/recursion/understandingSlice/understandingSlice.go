package main

import "fmt"

func len(v []int, index int, length int) {
	if index < 3 {
		fmt.Println(v)
		index++
		if index == length-1 {

			fmt.Println(v[1:index])
		} else {
			len(v[index:length], index, length)
		}

	}
}

func main() {
	v := []int{1, 2, 3}
	len(v, 0, 3)

}
