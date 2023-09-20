package main

import "fmt"

func info(name, msg string, age int) {
	// name := "Gopher"
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
	fmt.Printf("I'm %d year old\n", age)
}

func today() string {
	return "now"
}

// return 2 ตัว ใส่วงเล็บด้วบย
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	info("Kiza", "gopher", 111)
	fmt.Println()
	info("Go", "Lang", 14)
	fmt.Println()

	day := today()
	fmt.Println(day)

	a, b := swap("ใจ", "ดี")
	fmt.Println(a, b)

}
