package main

import (
	"errors"
	"fmt"
	"log"
)

func ReadFile(name string) (string, error) {
	// read file...
	// return "data...", nil

	return "", errors.New("file not found")
}

func main() {
	data, err := ReadFile("profile.txt")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("read file success : ", data)
}
