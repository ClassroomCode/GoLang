package main

import "fmt"

func main() {
	print(1)
}

func print(s interface{}) {
	st := s.(fmt.Stringer)
	fmt.Println(st.String())
}
