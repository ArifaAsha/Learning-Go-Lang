package main

import "fmt"

func printPrime(primes []int, length int, index int) {
	if index > 4 {
		return
	} else {
		if index == length-1 {
			fmt
			fmt.Println(primes[1:index])
		}
		fmt.Println(index)
		index++
		// var v []int = primes[index:5]
		printPrime(primes[index:4], 4, index)
	}

}

func main() {
	primes := []int{2, 3, 5, 7}
	// fmt.Println(len(primes))

	// var k []int = primes[0:6]
	// var s []int = primes[1:6]
	// var t []int = primes[2:6]
	// var u []int = primes[3:6]
	// var v []int = primes[4:6]

	// var w []int = primes[5:6]
	// fmt.Println(k)
	// fmt.Println(s)
	// fmt.Println(t)
	// fmt.Println(u)
	// fmt.Println(v)
	// fmt.Println(w)
	printPrime(primes, 4, 0)
}
