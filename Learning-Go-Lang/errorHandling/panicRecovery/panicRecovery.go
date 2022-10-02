package main

import "fmt"

func main() {
	// fmt.Println("Panic! in the Go app") //1
	// defer func() {                      //3
	// 	fmt.Println("I am differed func")
	// }()
	// panic("Oh no")             //2
	// fmt.Println("End of main") //unrechable

	fmt.Println("Help! Something bad is happening.")
	panic("Ending the program")

	fmt.Println("Waiting to execute")
}
