// section: code
package main

import (
	"fmt"
	"training/models"
)

func main() {
	a := models.DefaultAdmin()
	a.Greet()           // Call the `Greet` method that was promoted from the unexported user type
	fmt.Println(a.Name) // All exported/promoted fields are also accessible.
	// a.user.Greet() // this will not compile as `user` is not an exported field
}

// section: code

/*
// section: output
Hello, Homer
// section: output
*/
