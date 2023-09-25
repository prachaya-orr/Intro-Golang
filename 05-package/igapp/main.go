// สร้าง go mod init [pkg_module_name]
package main

import (
	"fmt"

	"github.com/prachaya/igapp/time"
	"github.com/prachaya/igapp/user"
)

func main() {
	user.Info("ohm", "gopher", 15)
	t := time.Today()
	fmt.Println("today is ", t)
}
