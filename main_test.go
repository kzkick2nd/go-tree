package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	// Create test dir
	tmpDir, err := prepareTestDirTree([]string{"test-dir/dir1", "test-dir/file1"})
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	code := m.Run()

	os.Exit(code)
}

func TestStdIn(t *testing.T) {
	// Parsing StdIn test.
}

func TestRun(t *testing.T) {
	// Unit test.
	// OK: Create sample directory and files before testing.
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
	// OK: type Tree(with leaf) type required.
	// OK: Create sample directory and files before testing.
	// TODO: Add option -L (length).

	cases := map[string]struct {
		src    string
		expect tree
	}{
		"1": {
			src: "./test-dir",
			expect: tree{v: "./test-dir", n: []tree{
				tree{v: "dir1"},
				tree{v: "file1"},
			}},
		},
		"2": {
			src: "test-dir",
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
func prepareTestDirTree(tree []string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("Error creating temp directory: %v", err)
	}

	for _, t := range tree {
		err = os.MkdirAll(filepath.Join(tmpDir, t), 0755)
		if err != nil {
			os.RemoveAll(tmpDir)
			return "", err
		}
	}

	return tmpDir, nil
}
