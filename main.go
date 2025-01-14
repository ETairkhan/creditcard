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
		if len(args) < 4 {
			fmt.Println("Usage: information --brands=<brand_file> --issuers=<issuer_file> <card_number>")
			os.Exit(1)
		}

		var brandFile, issuerFile, cardNumber string

		// Parse arguments
		for _, arg := range args[1:] {
			if strings.HasPrefix(arg, "--brands=") {
				brandFile = strings.TrimPrefix(arg, "--brands=")
			} else if strings.HasPrefix(arg, "--issuers=") {
				issuerFile = strings.TrimPrefix(arg, "--issuers=")
			} else {
				cardNumber = arg
			}
		}

		// Validate 
		if brandFile == "" || issuerFile == "" || cardNumber == "" {
			fmt.Println("Usage: information --brands=<brand_file> --issuers=<issuer_file> <card_number>")
			os.Exit(1)
		}

		
		brands := information.LoadData(brandFile)
		issuers := information.LoadData(issuerFile)
		information.CardInformation(cardNumber, brands, issuers)
	case "issue":
		
		if len(args) < 5 {
			fmt.Println("Usage: issue <brand_file> <issuer_file> <brand> <issuer>")
			os.Exit(1)
		}
		var brandFile, issuerFile, brand, issuer string

		for _, arg := range args[1:] {
			if strings.HasPrefix(arg, "--brands=") {
				brandFile = strings.TrimPrefix(arg, "--brands=")
			} else if strings.HasPrefix(arg, "--issuers=") {
				issuerFile = strings.TrimPrefix(arg, "--issuers=")
			} else if strings.HasPrefix(arg, "--brand=") {
				brand = strings.TrimPrefix(arg, "--brand=")
			} else if strings.HasPrefix(arg, "--issuer=") {
				issuer = strings.TrimPrefix(arg, "--issuer=")
			}
		}

		
		if brandFile == "" || issuerFile == "" || brand == "" || issuer == "" {
			fmt.Println("Usage: issue <brand_file> <issuer_file> <brand> <issuer>")
			os.Exit(1)
		}

		
		brands := information.LoadData(brandFile)
		issuers := information.LoadData(issuerFile)

		issue.IssuerCard(brands, issuers, brand, issuer)

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
