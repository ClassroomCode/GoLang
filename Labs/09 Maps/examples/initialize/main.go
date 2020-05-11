package main

import "fmt"

func main() {
	short()
	initialize()
	bad()

}

func short() {
	// section: short
	emails := map[string]string{
		"cory@gopherguides.com": "Cory LaNou",
		"mark@gopherguides.com": "Mark Bates",
		"tim@gopherguides.com":  "Tim Raymond",
	}
	// section: short

	fmt.Println(emails)

}

func initialize() {
	// section: make
	var emails map[string]string
	emails = make(map[string]string)
	emails["cory@gopherguides.com"] = "Cory LaNou"
	emails["mark@gopherguides.com"] = "Mark Bates"
	emails["tim@gopherguides.com"] = "Tim Raymond"
	// section: make

	fmt.Println(emails)

}

func bad() {
	// section: bad
	var emails map[string]string
	emails["cory@gopherguides.com"] = "Cory LaNou"
	// section: bad

	fmt.Println(emails)

}
