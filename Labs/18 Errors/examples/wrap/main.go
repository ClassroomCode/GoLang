package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// section: code
func main() {
	c, err := openConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}

type config struct {
	Backup bool `json:"backup"`
	Debug  bool `json:"debug"`
}

func openConfig(path string) (*config, error) {
	f, err := os.Open(path)
	if err != nil {
		// Wrap the error in our own message for the future
		return nil, fmt.Errorf("couldn't open config file: %w", err)
	}

	c := config{}

	if err := json.NewDecoder(f).Decode(&c); err != nil {
		log.Fatal(err)
	}
	return &c, nil
}

// section: code

/*
// section: output
2020/04/18 13:00:39 couldn't open config file: open ./config.json: no such file or directory
// section: output
*/
