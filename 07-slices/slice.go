package main

import "fmt"

func main() {
	// slice เหมือน array แต่ไม่ระบุขนาด
	langs := []string{"golang", "python", "javascript"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0]
	fmt.Printf("langs[0]: %#v\n", n)

	langs[1] = "pythonistas"
	fmt.Printf("langs: %#v\n", langs)

	// build-in function to find length of slice
	l := len(langs)
	fmt.Println("length of langs :", l)

	langs = append(langs, "F#", "Em", "Am")
	fmt.Printf("langs: %#v\n", langs)

	fmt.Println("length of langs :", len(langs))

}
