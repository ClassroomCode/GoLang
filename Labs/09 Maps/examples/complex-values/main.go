package main

import "fmt"

func main() {
	// section: code
	type Beatle struct {
		Name       string
		Instrument string
	}

	beatles := make(map[string]Beatle)

	// Create Beatle dynamically
	beatles["John"] = Beatle{Name: "John", Instrument: "guitar"}

	// Create Beatle from an instance
	b := Beatle{Name: "Paul", Instrument: "bass"}
	beatles[b.Name] = b

	b = Beatle{Name: "George"}
	beatles[b.Name] = b

	b = Beatle{Name: "Ringo", Instrument: "Drums"}
	beatles[b.Name] = b

	// Fix George:

	// Lookup George
	b, ok := beatles["George"]
	if !ok {
		fmt.Println("couldn't find George...")
		return
	}

	// Update George
	b.Instrument = "guitar"

	// Update the map with the new value
	beatles[b.Name] = b

	fmt.Println(beatles)
	// section: code
}

/*
// section: output
map[George:{George guitar} John:{John guitar} Paul:{Paul bass} Ringo:{Ringo Drums}]
// section: output

*/
