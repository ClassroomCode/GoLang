package main

import "time"

func main() {
	tick := time.Tick(100 * time.Millisecond)
	done := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			print("tick")
		case <-done:
			print(".done")
			return
		default:
			print(".")
			time.Sleep(10 * time.Millisecond)
		}
	}
}
