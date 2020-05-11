package main

import "fmt"

type Entertainer interface {
	Play()
}

type Beatle struct {
	Name       string
	Instrument string
}

func (b Beatle) Play() {
	fmt.Printf("%s plays %s", b.Name, b.Instrument)
}

func main() {
	b := Beatle{Name: "Ringo", Instrument: "Drums"}
	entertain(b)
}

func entertain(e Entertainer) {
	e.Play()
}
