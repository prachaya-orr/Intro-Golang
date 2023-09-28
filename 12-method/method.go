package main

import "fmt"

// ชื่อ variable ใน type struct ถ้า ตัวเล็กเป็น pubilc
// ตัวใหญ่ เป็น private

type User struct {
	Username      string
	Fullname      string
	ProfilePicUrl string
}

func (u User) info() {
	fmt.Printf("User name : %v\n", u.Username)
	fmt.Printf("Full name : %v\n", u.Fullname)
	fmt.Printf("Profile Picture URL : %v\n", u.ProfilePicUrl)
}
func main() {
	u := User{
		Username:      "golang",
		Fullname:      "Basic Golang",
		ProfilePicUrl: "https://knowhere.fake/gopher.jpg",
	}

	// info(u)
	u.info()
}
