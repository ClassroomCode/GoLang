package main

import "fmt"

func main() {
	m := map[string]string{
		"foo": "bar",
		"bin": "baz",
		"boo": "woo",
	}

	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	fmt.Printf("%+v", keys)
}
