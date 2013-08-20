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

	fmt.Println("all third level nodes with no key\n\n")
	fmt.Println(conf.Query().All("*").All("*").All(""))

	fmt.Println("last two include nodes under http\n\n")
	fmt.Println(conf.Query().All("http").LastN("include",2))

	fmt.Println("second top level node\n\n")
	fmt.Println(conf.Query().At("*", 1))

	fmt.Println("fourth to last node in first (happens to be the only) http\n\n")
	fmt.Println(conf.Query().First("http").At("*", -4))

	fmt.Println("all nodes in first server node in mail\n\n")
	fmt.Println(conf.Query().First("mail").First("server").All("*"))
}
