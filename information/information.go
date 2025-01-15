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
			fmt.Println("Card Band: -")
			fmt.Println("Card Issuer: -")
			return
		}

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
		fmt.Println()
	}
}
