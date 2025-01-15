package main

import (
	"creditcard/functional"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) <= 1 {
		functional.printUsage()
		os.Exit(1)
	}
	functional.Functional(args[1:])
}
