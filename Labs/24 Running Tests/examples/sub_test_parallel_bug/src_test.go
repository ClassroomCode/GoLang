package src

import "testing"

// section: test
func Test_Print(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "print 1", value: 1},
		{name: "print 2", value: 2},
		{name: "print 3", value: 3},
		{name: "print 4", value: 4},
	}
	for _, test := range tests {
		t.Run(test.name, func(st *testing.T) {
			st.Parallel()
			st.Log(test.value)
		})
	}
}

// section: test

/*
// section: output
$ go test -v ./sub_test_parallel_bug_test.go
=== RUN   Test_Print
=== PAUSE Test_Print
=== CONT  Test_Print
=== RUN   Test_Print/print_1
=== PAUSE Test_Print/print_1
=== RUN   Test_Print/print_2
=== PAUSE Test_Print/print_2
=== RUN   Test_Print/print_3
=== PAUSE Test_Print/print_3
=== RUN   Test_Print/print_4
=== PAUSE Test_Print/print_4
=== CONT  Test_Print/print_1
=== CONT  Test_Print/print_3
=== CONT  Test_Print/print_4
=== CONT  Test_Print/print_2
--- PASS: Test_Print (0.00s)
    --- PASS: Test_Print/print_1 (0.00s)
        sub_test_parallel_bug_test.go:20: 4
    --- PASS: Test_Print/print_3 (0.00s)
        sub_test_parallel_bug_test.go:20: 4
    --- PASS: Test_Print/print_4 (0.00s)
        sub_test_parallel_bug_test.go:20: 4
    --- PASS: Test_Print/print_2 (0.00s)
        sub_test_parallel_bug_test.go:20: 4
PASS
ok      command-line-arguments  (cached)
// section: output
*/
