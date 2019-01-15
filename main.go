package main

import "fmt"

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
	var t tree
	return t
}

func style(input tree) string {
	s := input.v
	return s
}
