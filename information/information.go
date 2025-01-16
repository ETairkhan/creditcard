package information

import (
	"bufio"
	"creditcard/validate"
	"fmt"
	"os"
	"strings"
)

func LoadData(filename string) map[string]string {
	data := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
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

func CardInformation(brands, issuers map[string]string, cards []string) {
	for _, card := range cards {
		valid := validate.IsValidLuhn(card)
		fmt.Println(card)
		if valid {
			fmt.Println("Correct: yes")
		} else {

			fmt.Println("Correct: no")
			fmt.Println("Card Brand: -")
			fmt.Println("Card Issuer: -")
			os.Exit(1)
			return
		}
		brand_result := "-"
		issuer_result := "-"
		for prefix, brand := range brands {
			if strings.HasPrefix(card, prefix) {
				brand_result = brand
				break
			}
		}
		fmt.Println("Card Brand:", brand_result)
		for prefix, issuer := range issuers {
			if strings.HasPrefix(card, prefix) {
				issuer_result = issuer
				break
			}
		}
		fmt.Println("Card Issuer:", issuer_result)
		fmt.Println()
	}
}
