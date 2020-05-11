package main

import (
	"fmt"

	"github.com/gopherguides/foo"
)

func main() {
	user := foo.NewUser("Cory", "LaNou", "p@ss0rd")
	// You can see the contents of the private information...
	fmt.Printf("%+v\n", user)
	// output:  {First:Cory Last:LaNou password:p@ss0rd}

	// You can't access or change it directly....

	//fmt.Println(user.password)
	//user.password = "new"

	// output:
	// ./foo_example.go:17:18: user.password undefined (cannot refer to unexported field or method password)
	// ./foo_example.go:18:6: user.password undefined (cannot refer to unexported field or method password)
}
