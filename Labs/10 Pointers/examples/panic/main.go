package main

func main() {
	// section: main
	var s *string
	*s = "boom"
	// section: main
}

/*
// section: output
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x104ea8d]

goroutine 1 [running]:
main.main()
        /Users/admin/go/src/github.com/gopherguides/learn/_training/fundamentals/pointers/src/panic.go:6 +0x1d
exit status 2
*/
