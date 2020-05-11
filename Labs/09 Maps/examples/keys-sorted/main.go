package main

import (
	"fmt"
	"sort"
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

	keys := make([]int, 0, len(months))

	for k := range months {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("keys: %+v\n", keys)
	for _, k := range keys {
		fmt.Println(k, months[k])
	}
}

// section: code

/*
// section: output
keys: [1 2 3 4 5 6 7 8 9 10 11 12]
1 January
2 February
3 March
4 April
5 May
6 June
7 July
8 August
9 September
10 October
11 November
12 December
// section: output
*/
