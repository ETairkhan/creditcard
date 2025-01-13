package main

import (
	"bufio"
	"creditcard/generate"
	"creditcard/validate"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "validate" {
		if len(args) > 1 && args[1] == "--stdin" {
			file, err := os.Open(args[2])
			if err != nil {
				panic(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				validate.Validate(strings.Fields(scanner.Text()))
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
		} else {
			validate.Validate(args[1:])
		}
	} else if len(args) > 0 && args[0] == "generate" {
		// Generation logic
		if len(args) < 2 {
			fmt.Println("Usage: generate <card_number_pattern> [--pick]")
			os.Exit(1)
		}

		// Detect --pick flag
		pick := false
		cardPattern := args[1]
		if len(args) > 2 && args[1] == "--pick" {
			pick = true
			cardPattern = args[2]
		}

		// Call the generate function
		generate.GenerateNumbers(cardPattern, pick)
	}
}
