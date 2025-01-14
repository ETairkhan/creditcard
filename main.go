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
			fmt.Println("1")
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
			fmt.Println("1")
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
			fmt.Println("1")
			os.Exit(1)
		}

		brands := information.LoadData(brandFile)
		issuers := information.LoadData(issuerFile)
		information.CardInformation(cardNumber, brands, issuers)
	case "issue":

		if len(args) < 5 {
			fmt.Println("1")
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
			fmt.Println("1")
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
	fmt.Println("  generate [--pick] <card_number_pattern>")
	fmt.Println("  information <brand_file> <issuer_file> <card_number>")
	fmt.Println("  issue <brand_file> <issuer_file> <brand> <issuer>")
}



// ./creditcard validat
// incorrect argument "validat"
// типа error message

// ./creditcard validate 3717631511358133831515213527517513715376135765312

// очень огромный номер (exit 1)
// сам я не считаю что это ошибка


// `./creditcard information --brands=issuers.txt --issuers=brands.txt "4400430180300003"`

// Exit code 1