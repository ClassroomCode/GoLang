package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

// section: code
var errConfigNotFound = errors.New("config not found")

func main() {
	c, err := openConfig("./config.json")
	if err != nil {
		if err == errConfigNotFound {
			fmt.Println("using default config")
			// ... load default config, etc.
			return
		}
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
		// Return our own sentinel error
		return nil, errConfigNotFound
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
using default config
// section: output
*/
