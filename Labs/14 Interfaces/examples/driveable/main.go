package main

import (
	"errors"
	"fmt"
	"log"
)

type driveable interface {
	Drive(int) error
}

func haveFun(d driveable) error {
	return d.Drive(55)
}

type Car struct{}

func (c Car) Drive(speed int) error {
	fmt.Printf("driving at %d\n", speed)
	return nil
}

type Bike struct{}

func (b Bike) Drive(speed int) error {
	if speed > 10 {
		return errors.New("can't drive that fast!")
	}
	fmt.Printf("driving at %d\n", speed)
	return nil
}

func main() {
	if err := haveFun(Car{}); err != nil {
		log.Fatal(err)
	}
	if err := haveFun(Bike{}); err != nil {
		log.Fatal(err)
	}
}
