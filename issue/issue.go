package issue

import (
	"creditcard/validate"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func IssuerCard(brands, issuers map[string]string, brand, issuer string) {

	brand = strings.TrimSpace(brand)
	issuer = strings.TrimSpace(issuer)

	brandPrefix := ""
	for prefix, b := range brands {
		if b == brand {
			brandPrefix = prefix
			break
		}
	}

	if brandPrefix == "" {
		fmt.Println("1")
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
		fmt.Println("1")
		os.Exit(1)
	}

	if !strings.HasPrefix(issuerPrefix, brandPrefix) {
		fmt.Println("1")
		os.Exit(1)
	}

	for {
		number := issuerPrefix
		for len(number) < 15 {
			number += fmt.Sprintf("%d", rand.Intn(10))
		}
		if validate.IsValidLuhn(number) {
			fmt.Println(number)
			break
		}
	}
}
