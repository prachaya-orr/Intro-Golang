package main

import "fmt"

// zero value of slice is nil
func main() {
	var langs []string
	// langs := []string{}
	fmt.Printf("langs : %#v\n", langs)

	if langs == nil {
		fmt.Println(`Yes nil "langs" is a nil slices`)
	} else {
		fmt.Printf("langs is NOT nil, value: %#v\n", langs)
	}
}
