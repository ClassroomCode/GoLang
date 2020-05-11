package main

import (
	"errors"
	"fmt"
	"log"
)

type errNotFound struct {
	key string
}

func (e *errNotFound) Error() string {
	return fmt.Sprintf("couldn't find value for key: %s", e.key)
}

func main() {
	s := NewStore()

	// check to see if key exists
	_, err := s.Get("Cory")
	if err != nil {
		var nf *errNotFound
		if errors.As(err, &nf) {
			s.Set("Cory", "cory@gopherguides.com")
		} else {
			log.Fatal(err)
		}
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
	return "", &errNotFound{key: k}
}
