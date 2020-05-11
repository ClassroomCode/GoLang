package main

import (
	"fmt"
	"log"
)

// #TODO: Declare a struct called `errNotFound`
// - Add an unexported field called `key` of type `string

// #TODO: Implement the error interface for `errNotFound`
// - It should return a string formatted as:
//   "couldn't find value for key: <key>"
//   hint: you can use the fmt.Sprintf function

func main() {
	s := NewStore()

	// check to see if key exists
	_, err := s.Get("Cory")
	// #TODO: check to see if this is a `errNotFound` error, and only if it is set the value
	// otherwise, log.Fatal the error
	if err != nil {
		s.Set("Cory", "cory@gopherguides.com")
	}

	v, err := s.Get("Cory")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}

type Store struct {
	keys map[string]string
}

func NewStore() *Store {
	return &Store{
		keys: make(map[string]string),
	}
}

func (s *Store) Set(k, v string) {
	s.keys[k] = v
}

func (s *Store) Get(k string) (string, error) {
	if v, ok := s.keys[k]; ok {
		return v, nil
	}
	// #TODO: return a errNotFound error instead of this generic error
	return "", fmt.Errorf("couldn't find value for key: %s", k)
}
