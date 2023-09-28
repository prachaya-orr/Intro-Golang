package main

import "fmt"

type User struct {
	Username      string
	Fullname      string
	ProfilePicUrl string
}

func main() {
	username := "golang"
	fullname := "Basic Golang"
	profilePicUrl := "https://knowhere.fake/gopher.jpg"
	fmt.Println(username, fullname, profilePicUrl)

	u := User{
		Username:      username,
		Fullname:      fullname,
		ProfilePicUrl: profilePicUrl,
	}
	// u.Username = username
	// u.Fullname = fullname
	// u.ProfilePicUrl = profilePicUrl

	fmt.Printf("%#v\n", u)

	name := u.Username
	fmt.Println(name)
	fmt.Println(u.Fullname)
	fmt.Println(u.ProfilePicUrl)
}
