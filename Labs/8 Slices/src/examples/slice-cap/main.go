package main

import "fmt"

func main() {
	// section: code
	names := []string{}
	fmt.Println("len:", len(names)) // 0
	fmt.Println("cap:", cap(names)) // 0

	names = append(names, "John")
	fmt.Println("len:", len(names)) // 1
	fmt.Println("cap:", cap(names)) // 1

	names = append(names, "Paul")
	fmt.Println("len:", len(names)) // 2
	fmt.Println("cap:", cap(names)) // 2

	names = append(names, "George")
	fmt.Println("len:", len(names)) // 3
	fmt.Println("cap:", cap(names)) // 4

	names = append(names, "Ringo")
	fmt.Println("len:", len(names)) // 4
	fmt.Println("cap:", cap(names)) // 4

	names = append(names, "Stu")
	// section: len
	fmt.Println("len:", len(names)) // 5
	fmt.Println("cap:", cap(names)) // 8
	// section: len
	// section: code
}
