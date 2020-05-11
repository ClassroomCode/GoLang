package main

import "log"

// section: code
type strError string

func (s strError) Error() string {
	return string(s)
}

func main() {
	err := boom()
	if err != nil {
		log.Fatal(err)
	}
}

func boom() error {
	return strError("something went wrong")
}

// section: code

/*
// section: output
2020/04/18 13:55:40 something went wrong
exit status 1
// section: output
*/
