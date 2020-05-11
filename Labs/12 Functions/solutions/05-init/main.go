package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var version = "0.1"
var name string

func init() {
	name = os.Args[0]
	if strings.Contains(name, string(os.PathSeparator)) {
		s := strings.Split(name, string(os.PathSeparator))
		name = s[len(s)-1]
	}
}

func main() {
	fmt.Printf("%s version %s starting at %s", name, version, time.Now().Format(time.Kitchen))
}
