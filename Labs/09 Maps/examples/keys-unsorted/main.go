package main

import (
	"fmt"
)

// section: code
func main() {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	for k, v := range months {
		fmt.Println(k, v)
	}

}

// section: code

/*
// section: output
$ go run ./keys-unsorted.go
6 June
8 August
12 December
9 September
10 October
1 January
2 February
3 March
4 April
5 May
7 July
11 November
// section: output
*/
