package information

import (
	"bufio"
	"creditcard/validate"
	"fmt"
	"os"
	"strings"
)

func loadData(filename string) map[string]string {
	data := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) == 2 {
			data[parts[1]] = parts[0]
		}
	}
	return data
}

func cardInformation(card string, brands, issuers map[string]string) {
	valid := validate.IsValidLuhn(card)
	fmt.Println("Card:", card)
	fmt.Println("Correct:", valid)
	for prefix, brand := range brands {
		if strings.HasPrefix(card, prefix) {
			fmt.Println("Card Brand:", brand)
			break
		}
	}
	for prefix, issuer := range issuers {
		if strings.HasPrefix(card, prefix) {
			fmt.Println("Card Issuer:", issuer)
			break
		}
	}
}
