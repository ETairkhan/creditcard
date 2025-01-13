package main

import (
	"bufio"
	"creditcard/generate"
	"creditcard/information"
	"creditcard/issue"
	"creditcard/validate"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}
	switch args[0] {
	case "validate":
		if len(args) > 1 && args[1] == "--stdin" {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				validate.Validate(strings.Fields(scanner.Text()))
			}

		} else {
			validate.Validate(args[1:])
		}
	case "generate":
		if len(args) < 2 {
			fmt.Println("Usage: generate <card_number_pattern> [--pick]")
			os.Exit(1)
		}

		pick := false
		cardPattern := args[1]
		if len(args) > 2 && args[1] == "--pick" {
			pick = true
			cardPattern = args[2]
		}

		generate.GenerateNumbers(cardPattern, pick)
	case "information":
		// Card information logic
		if len(args) < 4 {
			fmt.Println("Usage: info <card_number> <brand_file> <issuer_file>")
			os.Exit(1)
		}

		cardNumber := args[1]
		brandFile := args[2]
		issuerFile := args[3]

		brands := information.LoadData(brandFile)
		issuers := information.LoadData(issuerFile)
		information.CardInformation(cardNumber, brands, issuers)
	case "issue":
		// Card issuance logic
		if len(args) < 5 {
			fmt.Println("Usage: issue <brand> <issuer> <brand_file> <issuer_file>")
			os.Exit(1)
		}

		brand := args[1]
		issuer := args[2]
		brandFile := args[3]
		issuerFile := args[4]

		brands := information.LoadData(brandFile)
		issuers := information.LoadData(issuerFile)
		issue.IssueCard(brand, issuer, brands, issuers)

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  validate <card_numbers>")
	fmt.Println("  generate <card_number_pattern> [--pick]")
	fmt.Println("  information <card_number> <brand_file> <issuer_file>")
	fmt.Println("  issue <brand> <issuer> <brand_file> <issuer_file>")
}
