package main

import (
	"errors"
	"fmt"
)

// Messenger ensures we are a `message`
type Messenger interface {
	Message()
}

type Ack struct {
	Data []byte
}

func (a Ack) Message() {}
func (a Ack) Ack() string {
	return string(a.Data)
}

type Syn struct {
	Data []byte
}

func (s Syn) Message() {}
func (s Syn) Syn() string {
	return string(s.Data)
}

// This now uses the `messenger` interface, which is more desirable.
func processMessages(msg Messenger) error {
	switch m := msg.(type) {
	case Ack:
		fmt.Println("processing Ack message: ", m.Ack())
		return nil
	case Syn:
		fmt.Println("processingSyn message: ", m.Syn())
		return nil
	default:
		return errors.New("unknown message type!!!!")
	}
}

func main() {
	a := Ack{Data: []byte("ack message")}
	s := Syn{Data: []byte("syn message")}
	if err := processMessages(a); err != nil {
		fmt.Println(err)
	}
	if err := processMessages(s); err != nil {
		fmt.Println(err)
	}
	// This is no longer possible
	if err := processMessages("foo"); err != nil {
		fmt.Println(err)
	}

}
