package main

import "fmt"

func main() {
	// section: main

	s := new(string)
	i := new(int)

	type foo struct {
		id   int
		name string
	}
	f := new(foo)
	// Functionally equivalent to this:
	f1 := &foo{} // this is the idiomatic preferred style now



	fmt.Printf("s: %v, *s: %q\n", s, *s)
	fmt.Printf("i: %v, *i: %d\n", i, *i)
	fmt.Printf("f: %v, *f: %+v\n", f, *f)
	fmt.Printf("f1: %v, *f1: %+v\n", f1, *f1)

	// section: main
}

/*
// section: output
s: 0xc000092030, *s: ""
i: 0xc000094000, *i: 0
f: &{0 }, *f: {id:0 name:}
// section: output
*/
