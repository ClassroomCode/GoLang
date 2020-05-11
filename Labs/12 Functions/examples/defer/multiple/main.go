package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	b, err := loadPage()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))
}

// section: loadpage
func loadPage() ([]byte, error) {
	resp, err := http.Get("https://www.gopherguides.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // <<<
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %d", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// section: loadpage
