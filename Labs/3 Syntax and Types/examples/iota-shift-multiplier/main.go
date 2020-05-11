package main

import "fmt"

const (
	_  = 1 << (iota * 10) // ignore the first value
	KB                    // decimal:       1024 -> binary 00000000000000000000010000000000
	MB                    // decimal:    1048576 -> binary 00000000000100000000000000000000
	GB                    // decimal: 1073741824 -> binary 01000000000000000000000000000000
)

func main() {
	fmt.Printf("KB =  %v\n", KB)
	fmt.Printf("MB =  %v\n", MB)
	fmt.Printf("GB =  %v\n", GB)
}
