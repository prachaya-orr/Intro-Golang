package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	langs := [5]string{"golang", "python", "javascript", "F#"}
	fmt.Println("\nfor basic")

	for i := 0; i < len(langs); i++ {
		value := langs[i]
		fmt.Println(i, ":", value)
	}

	fmt.Println("\nfor basic")
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}

	fmt.Println("\nfor basic")
	for _, value := range langs {
		fmt.Println("only value :", value)
	}

	fmt.Println("\nfor basic")
	for index := range langs {
		fmt.Println("only index :", index)
	}

}
