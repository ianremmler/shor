package main

import (
	"github.com/ianremmler/shor"

	"flag"
	"fmt"
	"os"
)

func main() {
	isSingle, indent := false, "\t"
	flag.BoolVar(&isSingle, "s", isSingle, "Produce single-line output")
	flag.StringVar(&indent, "i", indent, "String used to indent one level")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: shor [-s] [-i indentstr] < input.shor")
		flag.PrintDefaults()
	}
	flag.Parse()

	tree, err := shor.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	depth := 0
	if isSingle {
		depth = -1
	}
	fmt.Print(tree.Format(depth, indent))
	if isSingle {
		fmt.Println()
	}
}
