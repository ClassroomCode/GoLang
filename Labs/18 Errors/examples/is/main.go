package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	var f io.Reader
	var err error
	f, err = os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// use default config
			f = bytes.NewBufferString(`{"backup":false,"debug":true}`)
		} else {
			return nil, fmt.Errorf("something went terribly wrong...: %w", err)
		}
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
&{Backup:false Debug:true}
// section: output
*/
