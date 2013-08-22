package main

import (
	"github.com/ianremmler/shor"

	"fmt"
	"os"
)

func main() {
	in, err := os.Open("nginx.shor")
	if err != nil {
		os.Exit(1)
	}
	defer in.Close()
	conf, err := shor.Parse(in)
	if err != nil {
		os.Exit(1)
	}

	// some random queries

	fmt.Println("all third level nodes with no key")
	fmt.Println(conf.Query().All("*").All("*").All(""))

	fmt.Println("\nlast two include nodes under http")
	fmt.Println(conf.Query().All("http").LastN("include", 2))

	fmt.Println("\nsecond top level node")
	fmt.Println(conf.Query().At("*", 1))

	fmt.Println("\nfourth to last node in first (happens to be the only) http")
	fmt.Println(conf.Query().First("http").At("*", -4))

	fmt.Println("\nall nodes in first server node in mail")
	fmt.Println(conf.Query().First("mail").First("server").All("*"))

	// get and set some values

	s, q := "", conf.Query().First("user")
	if len(q) > 0 && q[0].Get(&s) {
		fmt.Println("\nuser:", s)
	}
	n, q := 0.0, conf.Query().First("worker_processes")
	if len(q) > 0 && q[0].Get(&n) {
		fmt.Println("\nworker_processes:", n)
	}
	b, q := false, conf.Query().First("http").First("gzip")
	if len(q) > 0 && q[0].Get(&b) {
		fmt.Println("\ngzip:", b)
		q[0].Set(!b)
		q[0].Get(&b)
		fmt.Println("gzip set to opposite:", b)
	}
}
