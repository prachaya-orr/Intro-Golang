package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday
	// switch ใน Go ไม่ต้องใส่ break
	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
		fallthrough // ถ้าอยากให้้เข้า monday และตกไป case ข้างล่างด้วย
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}
}
