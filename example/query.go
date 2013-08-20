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
	fmt.Println(conf.Query().All("http").LastN("include",2))

	fmt.Println("\nsecond top level node")
	fmt.Println(conf.Query().At("*", 1))

	fmt.Println("\nfourth to last node in first (happens to be the only) http")
	fmt.Println(conf.Query().First("http").At("*", -4))

	fmt.Println("\nall nodes in first server node in mail")
	fmt.Println(conf.Query().First("mail").First("server").All("*"))
}
