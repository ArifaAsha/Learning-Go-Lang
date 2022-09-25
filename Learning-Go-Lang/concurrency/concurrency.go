package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("fish")
}

func count(s string) {
	for i := 1; i<5; i++ {
		fmt.Println(i, s)

		time.Sleep(time.Millisecond * 500)

	}
}
