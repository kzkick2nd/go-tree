package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type tree struct {
	v string
	n []tree
}

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	fmt.Println(run(dir))
}

func run(stdin string) string {
	t := seek(stdin)
	s := style(t)
	return s
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
	s := drawTree(t, 0, []bool{false})
	return s
}

func drawTree(t tree, indent int, eol []bool) string {
	var s string
	eol = append(eol, false)

	s += t.v + "\n"

	for i, l := range t.n {
		for j := 1; j <= indent; j++ {
			if !eol[j-1] {
				s += "|  "
			} else {
				s += "   "
			}
		}

		if i == len(t.n)-1 {
			s += "└── "
			eol[indent] = true
		} else {
			s += "├── "
			eol[indent] = false
		}

		s += drawTree(l, indent+1, eol)
	}
	return s
}
