package main

import "fmt"

type Beatle struct {
	Name       string
	Instrument string
}

func (b Beatle) String() string {
	return fmt.Sprintf("%s plays %s", b.Name, b.Instrument)
}

func main() {
	b := Beatle{Name: "Ringo", Instrument: "Drums"}
	fmt.Println(b.String()) // Ringo plays Drums
}
