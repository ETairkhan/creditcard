package issue

import (
	"creditcard/validate"
	"fmt"
	"math/rand"
	"os"
)

func IssuerCard(brands, issuers map[string]string, brand, issuer string) {
	prefix := ""
	for p, b := range brands {
		if b == brand {
			prefix = p
			break
		}
	}
	for p, i := range issuers {
		if i == issuer {
			prefix = p
			break
		}
	}

	if prefix == "" {
		fmt.Println("Error: Brand or Issuer not found")
		os.Exit(1)
	}

	for {
		number := prefix
		for len(number) < 15 {
			number += fmt.Sprintf("%d", rand.Intn(10))
		}
		if validate.IsValidLuhn(number) {
			fmt.Println(number)
			break
		}
	}
}
