package main

import "fmt"

// section: verbs
// Use %v to print the value
// Use %s to print a string
// Use %q to quote a string
// Use %d to print a decimal (int, int32, etc)
// Use %T to print out the data type of the variable
// Use \ to escape a character, like a quote:  \"
// Use \n to print a new line (line return)
// Use \t to insert a tab
// Use %+v to print the name and value
// section: verbs

func main() {

	// section: basic
	fmt.Println("Print a statement followed by a line return")
	fmt.Printf("Hello %s\n", "gopher")
	// section: basic

	// section: print
	a := 1
	fmt.Println("This will join all arguments and print them.  a =  ", a)

	fmt.Printf("This is the `format` string. Escape \" and print value of a: %v\n", a)

	type User struct {
		Name string
	}
	u := User{Name: "Mary"}
	fmt.Printf("user: %+v", u)
	// section: print

}

/*
// section: basic-output
Print a statement followed by a line return
Hello gopher
// section: basic-output
*/

/*
// section: verb-output
This will join all arguments and print them.  a =   1
This is the `format` string. Escape " and print value of a: 1
user: {Name:Mary}
// section: verb-output
*/
