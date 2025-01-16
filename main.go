package main

import (
	"creditcard/mainFunc"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		mainFunc.PrintUsage()
		os.Exit(1)
	}
	mainFunc.Functional(args)
}
