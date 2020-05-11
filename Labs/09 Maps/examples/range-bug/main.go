package main

import (
	"fmt"
	"strings"
)

func main() {
	cpy()
	key()
}

func cpy() {
	// section: copy
	type Beatle struct {
		Name       string
		Instrument string
	}

	beatles := make(map[string]Beatle)

	beatles["John"] = Beatle{Name: "John", Instrument: "guitar"}
	beatles["Paul"] = Beatle{Name: "Paul", Instrument: "bass"}
	beatles["George"] = Beatle{Name: "George", Instrument: "guitar"}
	beatles["Ringo"] = Beatle{Name: "Ringo", Instrument: "drums"}

	for _, b := range beatles {
		// `b` is a COPY, and will not update the value stored in the map
		b.Instrument = strings.ToUpper(b.Instrument)
		fmt.Println(b)
	}

	fmt.Println(beatles)
	// section: copy

	/*
		// section: copy-output
		{John GUITAR}
		{Paul BASS}
		{George GUITAR}
		{Ringo DRUMS}
		map[George:{George guitar} John:{John guitar} Paul:{Paul bass} Ringo:{Ringo drums}]
		// section: copy-output
	*/

}

func key() {
	// section: key
	type Beatle struct {
		Name       string
		Instrument string
	}

	beatles := make(map[string]Beatle)

	beatles["John"] = Beatle{Name: "John", Instrument: "guitar"}
	beatles["Paul"] = Beatle{Name: "Paul", Instrument: "bass"}
	beatles["George"] = Beatle{Name: "George", Instrument: "guitar"}
	beatles["Ringo"] = Beatle{Name: "Ringo", Instrument: "drums"}

	for k, b := range beatles {
		b.Instrument = strings.ToUpper(b.Instrument)
		fmt.Println(b)
		// update the map:
		beatles[k] = b
	}

	fmt.Println(beatles)
	// section: key

	/*
		// section: key-output
		{John GUITAR}
		{Paul BASS}
		{George GUITAR}
		{Ringo DRUMS}
		map[George:{George GUITAR} John:{John GUITAR} Paul:{Paul BASS} Ringo:{Ringo DRUMS}]
		// section: key-output
	*/
}
