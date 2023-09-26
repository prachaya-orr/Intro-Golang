package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}
	fmt.Printf("% #v\n", status[200])

	l := len(status)
	fmt.Printf("length: %d\n\n", l)

	status[200] = "Okie"
	status[285] = "GG EZ"

	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	value := status[200]
	fmt.Printf("%#v\n", value)

	delete(status, 285)
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	v, ok := status[666]
	if ok {
		fmt.Printf("value : %q \n\n", v) // empty value : ""
	} else {
		fmt.Println("not found")
	}

	if p, ok := status[666]; ok {
		fmt.Printf("value : %q \n\n", p) // empty value : ""
	} else {
		fmt.Println("not found")
	}

	// var m map[string]string = make(map[string]string) // จองที่ mem ให้แล้ว ไม่ nil แล้ว
	// var m = make(map[string]string)
	// m := make(map[string]string)
	// m := map[string]string{}
	m := map[string]string{
		"ออเจ้า": "+66-978788997",
		"Steven": "+61-857985555",
	}

	if m == nil {
		fmt.Printf("m is nil. value : %#v\n", m)
	} else {
		fmt.Printf("m is not nil. value : %#v\n", m)
	}
}
