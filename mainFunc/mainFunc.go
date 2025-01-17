package mainFunc

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

func Functional(args []string) {
	switch args[0] {
	case "validate":
		if len(args) < 2 {
			fmt.Println("Error: Missing card numbers or --stdin flag for validation")
			PrintUsage()
			os.Exit(1)
		}

		useStdin := false

		for _, arg := range args {
			if arg == "--stdin" {
				useStdin = true
				break
			}
		}

		if useStdin {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line == "exit" {
					fmt.Println("Exiting...")
					os.Exit(0)
				}
				if line == "" {

					fmt.Println("INCORRECT")
					os.Exit(1)

				}
				numbers := strings.Fields(line)
				if len(numbers) > 0 {
					validate.Validate(numbers)
				} else {

					fmt.Println("INCORRECT")
					os.Exit(1)
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
		} else {
			validate.Validate(args[1:])
		}

	case "generate":
		if len(args) < 2 {
			fmt.Println("Error: Missing card number pattern for 'generate'")
			fmt.Println("Usage: generate [--pick] <card_number_pattern>")
			os.Exit(1)
		}

		pick := false
		cardPattern := ""

		for _, arg := range args[1:] {
			if arg == "--pick" {
				pick = true
				continue
			}

			if cardPattern == "" && cardPattern != "generate" {
				cardPattern = arg
			} else {
				continue
			}
		}
		if cardPattern == "" {
			fmt.Println("Error: Missing card number pattern for 'generate'")
			fmt.Println("Usage: generate [--pick] <card_number_pattern>")
			os.Exit(1)
		}

		generate.GenerateNumbers(cardPattern, pick)
	case "information":
		if len(args) < 3 {
			fmt.Println("Error: Missing required arguments for 'information'")
			fmt.Println("Usage: information --brands=<file> --issuers=<file> <card_number>")
			os.Exit(1)
		}

		var brandFile, issuerFile string
		var cardNumbers []string
		useStdin := false

		for _, arg := range args[1:] {
			if strings.HasPrefix(arg, "--brands=") {
				brandFile = strings.TrimPrefix(arg, "--brands=")
				if brandFile != "brands.txt" {
					os.Exit(1)
				}
			} else if strings.HasPrefix(arg, "--issuers=") {
				issuerFile = strings.TrimPrefix(arg, "--issuers=")
				if issuerFile != "issuers.txt" {
					os.Exit(1)
				}
			} else if arg == "--stdin" {
				useStdin = true
			} else {
				cardNumbers = append(cardNumbers, arg)
			}
		}
		if brandFile == "" || issuerFile == "" {
			fmt.Println("Error: Missing required arguments for 'information'")
			fmt.Println("Usage: information --brands=<file> --issuers=<file> <card_number>")
			os.Exit(1)
		}

		brands := information.LoadData(brandFile)
		if brands == nil {
			fmt.Printf("Error: Could not load brands data from %s\n", brandFile)
			os.Exit(1)
		}
		validate.ValidateData(brands, "brands")

		issuers := information.LoadData(issuerFile)
		if issuers == nil {
			fmt.Printf("Error: Could not load issuers data from %s\n", issuerFile)
			os.Exit(1)
		}
		validate.ValidateData(issuers, "issuers")

		if useStdin {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" {
					fields := strings.Fields(line)
					cardNumbers = append(cardNumbers, fields...)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
		}

		if len(cardNumbers) == 0 {
			fmt.Println("Error: Missing required arguments for 'information'")
			fmt.Println("Usage: information --brands=<file> --issuers=<file> <card_number>")
			os.Exit(1)
		}

		information.CardInformation(brands, issuers, cardNumbers)
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

		brands := make(map[string][]string)
		brandData, err := os.ReadFile(brandFile)
		if err != nil {
			fmt.Printf("Error: Failed to read brands file '%s': %v\n", brandFile, err)
			os.Exit(1)
		}

		lines := strings.Split(string(brandData), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				fmt.Printf("Error: Invalid line in brands file: '%s'\n", line)
				os.Exit(1)
			}
			brandName := strings.TrimSpace(parts[0])
			prefix := strings.TrimSpace(parts[1])
			brands[brandName] = append(brands[brandName], prefix)
		}

		issuers := information.LoadData(issuerFile)
		validate.ValidateData(issuers, "issuers")

		issue.IssuerCard(brands, issuers, brand, issuer)

	default:
		PrintUsage()
		os.Exit(1)
	}
}

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  validate <card_numbers>          - Validate one or more card numbers")
	fmt.Println("  validate --stdin                 - Validate card numbers via standard input")
	fmt.Println("  generate [--pick] <pattern>      - Generate card numbers using a pattern")
	fmt.Println("  information --brands=<file> --issuers=<file> <card_number>")
	fmt.Println("                                   - Retrieve brand and issuer information")
	fmt.Println("  issue --brands=<file> --issuers=<file> --brand=<brand> --issuer=<issuer>")
	fmt.Println("                                   - Issue a new card with specified brand and issuer")
}
