package main

import (
	"fmt"
)

func main() {
	months := []string{
		"",
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	
	q2 := months[4:7]
	summer := months[6:9]
	
	q2 = append(q2, "T1", "T2", "T3", "T4", "T5", "T6") // "T7"
	
	fmt.Printf("%v %d|%d\n", months, len(months), cap(months))
	fmt.Printf("%v %d|%d\n", q2, len(q2), cap(q2))
	fmt.Printf("%v %d|%d\n", summer, len(summer), cap(summer))
}