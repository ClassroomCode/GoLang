package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// section: code
	type config struct {
		Backup bool `json:"backup"`
		Debug  bool `json:"debug"`
	}

	var f io.Reader
	var err error

	// try to read a file
	f, err = os.Open("./config.json")
	if err != nil {
		// Fall back to a default config. Stream it as a reader.
		f = strings.NewReader(`{"backup":false,"debug":true}`)
	}

	c := config{}

	if err := json.NewDecoder(f).Decode(&c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)
	// section: code
}

/*
// section: output
{Backup:false Debug:true}
// section: output
*/
