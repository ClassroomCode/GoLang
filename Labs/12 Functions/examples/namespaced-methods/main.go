package main

import "fmt"

type Beatle struct {
	Name       string
	Instrument string
}

// section: code
func (b Beatle) String() string {
	return fmt.Sprintf("%s plays %s", b.Name, b.Instrument)
}

func String(b Beatle) string {
	return fmt.Sprintf("%s plays %s", b.Name, b.Instrument)
}

func main() {
	b := Beatle{Name: "Ringo", Instrument: "Drums"}
	fmt.Println(b.String()) // Ringo plays Drums
	fmt.Println(String(b))  // Ringo plays Drums
}

// section: code
