package main

import (
	"github.com/ianremmler/thor"

	"fmt"
	"os"
)

func main() {
	tree, err := thor.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(tree)
}
