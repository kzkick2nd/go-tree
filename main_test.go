package main

import "testing"

func TestStdIn(t *testing.T) {
	// Parsing StdIn test.
}

func TestCmd(t *testing.T) {
	// Unit test.
	// TODO: Create sample directory and files before testing.
	cases := map[string]struct {
		stdin  string
		expect string
	}{
		"1": {stdin: "./test-dir", expect: "./test-dir\n├── dir1\n└── file1"},
	}
	for n, tc := range cases {
		tc := tc
		tc.Run(n, func(t *testing.T) {
			if got := tree(tc.stdin); got != tc.expect {
				t.Fatalf("expected is %v, but got is %V", tc.expext, got)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	// Test walk and find path function.
	// TODO: Create sample directory and files before testing.
	// TODO: Add option -L (length).
}

func TestStyle(t *testing.T) {
	// Test styling output from array.
	// TODO: Add option to set color output.
}
