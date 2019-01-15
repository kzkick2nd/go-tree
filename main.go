package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	var p string
	fmt.Scan(&p)
	fmt.Println(run(p))
}

func run(stdin string) string {
	t := seek(stdin)
	s := style(t)
	return s
}

type tree struct {
	v string
	n []tree
}

func seek(src string) tree {
	t := seekDir(src, src)
	return t
}

func seekDir(name, src string) tree {
	t := tree{v: name}

	list, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range list {
		if f.IsDir() {
			t.n = append(t.n, seekDir(f.Name(), filepath.Join(src, f.Name())))
			continue
		}
		t.n = append(t.n, tree{v: f.Name(), n: []tree{}})
	}
	return t
}

func style(t tree) string {
	s := drawTree(t, 0, false)
	return s
}

func drawTree(t tree, indent int, eol bool) string {
	var s string
	s += t.v + "\n"
	for i, l := range t.n {
		eol := eol

		if indent > 1 {
			s += "│"
		}

		for j := 1; j <= indent; j++ {
			if j == indent && !eol {
				s += "│  "
			} else {
				s += "   "
			}
		}

		if i == len(t.n)-1 {
			s += "└── "
			eol = true
		} else {
			s += "├── "
			eol = false
		}

		s += drawTree(l, indent+1, eol)
	}
	return s
}
