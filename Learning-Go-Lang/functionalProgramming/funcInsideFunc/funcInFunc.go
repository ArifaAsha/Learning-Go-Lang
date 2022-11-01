package main

import "fmt"

func main() {
	test := //store the returned value of the func
		func(x int) int { //dosen't have explicit name
			return x * -1
		}(8) //call the anonymous func with parameter 8
	fmt.Println(test)
}
