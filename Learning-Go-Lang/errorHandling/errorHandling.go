package main

import (
	"errors"
	"fmt"
	"log"
)

func doError() (string, error) { //error should be last
	return "", errors.New("This is an error") //creating new error
}

func doNoError() (string, error) {
	return "My response", nil
}

func main() {
	resp, err := doError()
	if err != nil {
		log.Printf("There was an error: %v\n", err)
	}
	fmt.Println("My message: ", resp)

	resp, err = doNoError()
	if err != nil {
		log.Printf("Should not print")
	}
	fmt.Println("My response: ", resp)
}
