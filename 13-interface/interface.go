package main

import (
	"fmt"
)

type Phone interface {
	Call(number string)
}

type Samsung struct {
	Name string
}

func (s Samsung) Call(number string) {
	fmt.Println(s.Name, "calling", number)
}

func (s Samsung) Answer() {
	fmt.Println(s.Name, "hello there!")
}

func Dial(p Phone) {
	p.Call("+66823424141")
}

func main() {
	s := Samsung{
		Name: "S10",
	}

	s.Answer()

	Dial(s)
}
