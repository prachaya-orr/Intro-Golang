package main

import "fmt"

func main() {
	num := 31

	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else if num == 31 {
		fmt.Println("เลยสวย 31")
	} else {
		fmt.Println("เลขคี่")
	}
}
