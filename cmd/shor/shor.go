package main

import (
	"github.com/ianremmler/shor"

	"fmt"
	"os"
)

func main() {
	tree, err := shor.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(tree)
}
