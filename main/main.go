package main

import (
	"bufio"
	"creditcard/validate"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "validate" {
		if len(args) > 1 && args[1] == "--stdin" {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				validate.Validate(strings.Fields(scanner.Text()))
			}
		} else {
			validate.Validate(args[1:])
		}
	}
}
