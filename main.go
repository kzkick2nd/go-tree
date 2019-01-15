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
	t := lsr(src, src)
	return t
}

func lsr(name, src string) tree {
	t := tree{v: name}

	list, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range list {
		if f.IsDir() {
			t.n = append(t.n, lsr(f.Name(), filepath.Join(src, f.Name())))
			continue
		}
		t.n = append(t.n, tree{v: f.Name(), n: []tree{}})
	}
	return t
}

func style(input tree) string {
	s := input.v
	return s
}
