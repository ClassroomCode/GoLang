package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var version = "0.1"

func main() {
	// #TODO: Move all the logic to parse name to the an `init` function
	var name = ""
	name = os.Args[0]
	if strings.Contains(name, string(os.PathSeparator)) {
		s := strings.Split(name, string(os.PathSeparator))
		name = s[len(s)-1]
	}
	fmt.Printf("%s version %s starting at %s", name, version, time.Now().Format(time.Kitchen))
}
