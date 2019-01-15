package main

import (
	"reflect"
	"testing"
)

func TestStdIn(t *testing.T) {
	// Parsing StdIn test.
}

func TestRun(t *testing.T) {
	// Unit test.
	// TODO: Create sample directory and files before testing.
	cases := map[string]struct {
		stdin  string
		expect string
	}{
		"1": {stdin: "./test-dir", expect: "./test-dir\n├── dir1\n└── file1"},
		"2": {stdin: "test-dir", expect: "test-dir\n├── dir1\n└── file1"},
	}
	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := run(tc.stdin); actual != tc.expect {
				t.Fatalf("expect is %v, but actual is %v", tc.expect, actual)
			}
		})
	}
}

func TestSeek(t *testing.T) {
	// Test walk and find path function.
	// type Tree(with leaf) type required.
	// TODO: Create sample directory and files before testing.
	// TODO: Add option -L (length).

	cases := map[string]struct {
		src    string
		expect tree
	}{
		"1": {
			src: "./test-dir",
			expect: tree{v: "test-dir", n: []tree{
				tree{v: "dir1"},
				tree{v: "file1"},
			}},
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := seek(tc.src); !reflect.DeepEqual(actual, tc.expect) {
				t.Fatalf("expect is %v, but actual is %v", tc.expect, actual)
			}
		})
	}
}

func TestStyle(t *testing.T) {
	// Test styling output from array.
	// TODO: Add option to set color output.
	cases := map[string]struct {
		input  tree
		expect string
	}{
		"1": {
			input: tree{v: "test-dir", n: []tree{
				tree{v: "dir1"},
				tree{v: "file1"},
			}},
			expect: "test-dir\n├── dir1\n└── file1",
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := style(tc.input); actual != tc.expect {
				t.Fatalf("expect is %v, but actual is %v", tc.expect, actual)
			}
		})
	}
}
