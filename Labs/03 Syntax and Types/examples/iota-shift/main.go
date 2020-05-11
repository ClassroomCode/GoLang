package main

import "fmt"

const (
	read   = 1 << iota // 00000001 = 1
	write              // 00000010 = 2
	remove             // 00000100 = 4

	// admin will have all of the permissions
	admin = read | write | remove
)

func main() {
	fmt.Printf("read =  %v\n", read)
	fmt.Printf("write =  %v\n", write)
	fmt.Printf("remove =  %v\n", remove)
	fmt.Printf("admin =  %v\n", admin)
}
