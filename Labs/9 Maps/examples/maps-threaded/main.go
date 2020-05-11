package main

func main() {
	numbers := map[int]int{}
	for i := 0; true; i++ {
		go func(i int) {
			numbers[i] = i
		}(i)
	}
}
