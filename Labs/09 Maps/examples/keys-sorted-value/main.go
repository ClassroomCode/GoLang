package main

import (
	"fmt"
	"sort"
)

// section: code
func main() {
	words := map[int]string{
		0:  "time",
		1:  "person",
		2:  "year",
		3:  "way",
		4:  "day",
		5:  "thing",
		6:  "man",
		7:  "world",
		8:  "life",
		9:  "hand",
		10: "part",
		12: "child",
		13: "eye",
		14: "woman",
		15: "place",
		16: "work",
		17: "week",
		18: "case",
		19: "point",
		20: "government",
		21: "company",
		22: "number",
		23: "group",
		24: "problem",
		25: "fact",
	}

	// create an inverted map where values are now the keys
	sorted := make(map[string]int)

	keys := make([]string, 0, len(words))

	for k, v := range words {
		keys = append(keys, v)
		sorted[v] = k
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(sorted[k], k)
	}
}

// section: code

/*
// section: output
18 case
12 child
21 company
4 day
13 eye
25 fact
20 government
23 group
9 hand
8 life
6 man
22 number
10 part
1 person
15 place
19 point
24 problem
5 thing
0 time
3 way
17 week
14 woman
16 work
7 world
2 year
// section: output
*/
