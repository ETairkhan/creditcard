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
			cardNumber := strings.Join(args[1:], "") // Join all parts into one number
			cardNumber = strings.ReplaceAll(cardNumber, " ", "") // Remove spaces
			validate.Validate([]string{cardNumber})

		}
	case "generate":
		if len(args) < 2 {
			fmt.Println("Error: Missing card number pattern for 'generate'")
			fmt.Println("Usage: generate [--pick] <card_number_pattern>")
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
			fmt.Println("Error: Missing required arguments for 'information'")
			fmt.Println("Usage: information --brands=<file> --issuers=<file> <card_number>")
			os.Exit(1)
		}

		var brandFile, issuerFile, cardNumber string

		// Parse arguments
		for _, arg := range args[1:] {
			if strings.HasPrefix(arg, "--brands=") {
				brandFile = strings.TrimPrefix(arg, "--brands=")
				if brandFile != "brands.txt" {
					os.Exit(1)
				}
			} else if strings.HasPrefix(arg, "--issuers=") {
				issuerFile = strings.TrimPrefix(arg, "--issuers=")
				if brandFile != "issuers.txt" {
					os.Exit(1)
				}
			} else {
				cardNumber = arg
			}
		}

		// Validate
		if brandFile == "" || issuerFile == "" || cardNumber == "" {
			fmt.Println("Error: Missing required arguments for 'information'")
			fmt.Println("Usage: information --brands=<file> --issuers=<file> <card_number>")
			os.Exit(1)
		}

		brands := information.LoadData(brandFile)
		validate.ValidateData(brands, "brands")

		issuers := information.LoadData(issuerFile)
		validate.ValidateData(issuers, "issuers")
		// Validate if card number matches any brand prefix

		information.CardInformation(cardNumber, brands, issuers)
	case "issue":

		if len(args) < 5 {
			fmt.Println("Error: Missing required arguments for 'issue'")
			fmt.Println("Usage: issue --brands=<file> --issuers=<file> --brand=<brand> --issuer=<issuer>")
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
			fmt.Println("Error: Missing required arguments for 'issue'")
			fmt.Println("Usage: issue --brands=<file> --issuers=<file> --brand=<brand> --issuer=<issuer>")
			os.Exit(1)
		}

		brands := information.LoadData(brandFile)
		validate.ValidateData(brands, "brands")

		issuers := information.LoadData(issuerFile)
		validate.ValidateData(issuers, "issuers")

		issue.IssuerCard(brands, issuers, brand, issuer)

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  validate <card_numbers>")
	fmt.Println("  generate [--pick] <card_number_pattern>")
	fmt.Println("  information <brand_file> <issuer_file> <card_number>")
	fmt.Println("  issue <brand_file> <issuer_file> <brand> <issuer>")
}
