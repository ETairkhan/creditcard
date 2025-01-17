package issue

import (
	"creditcard/validate"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func IssuerCard(brands map[string][]string, issuers map[string]string, brand, issuer string) {
	brand = strings.TrimSpace(brand)
	issuer = strings.TrimSpace(issuer)

	brandPrefixes, found := brands[brand]
	if !found {
		fmt.Println("Error: Brand not found in the provided data.")
		os.Exit(1)
	}

	issuerPrefix := ""
	for prefix, i := range issuers {
		if i == issuer {
			issuerPrefix = prefix
			break
		}
	}

	if issuerPrefix == "" {
		fmt.Println("Error: Issuer not found in the provided data.")
		os.Exit(1)
	}

	validPrefix := false
	for _, brandPrefix := range brandPrefixes {
		if strings.HasPrefix(issuerPrefix, brandPrefix) {
			validPrefix = true
			break
		}
	}

	if !validPrefix {
		fmt.Printf("Error: Issuer prefix '%s' does not match any prefix for brand '%s'.\n", issuerPrefix, brand)
		os.Exit(1)
	}

	maxAttempts := 10000
	for attempts := 0; attempts < maxAttempts; attempts++ {
		number := issuerPrefix
		for len(number) < 15 {
			number += fmt.Sprintf("%d", rand.Intn(10))
		}
		if validate.IsValidLuhn(number) {
			fmt.Println(number)
			return
		}
	}
	fmt.Println("Error: Failed to generate a valid Luhn number after maximum attempts.")
	os.Exit(1)
}
