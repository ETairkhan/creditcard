package generate

import (
	"creditcard/validate"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GenerateNumbers(card string, pick bool) {
	if len(card) < 13 || len(card) > 19 {
		fmt.Println("Error: not correct length for the card")
		os.Exit(1)
	}
	firstAsterisk := strings.Index(card, "*")
	if firstAsterisk != -1 && firstAsterisk != len(card)-strings.Count(card[firstAsterisk:], "*") {
		fmt.Println("Error: Asterisks must only appear at the end.")
		os.Exit(1)
	}

	if strings.Count(card, "*") > 4 || strings.Count(card, "*") == 0 {
		fmt.Println("Error: Many asterisk or no asterisk")
		os.Exit(1)
	}

	results := []string{}
	generateRecursive(card, "", &results)

	valid := []string{}
	for _, num := range results {
		if validate.IsValidLuhn(num) {
			valid = append(valid, num)
		} else {
			continue
		}
	}

	if pick && len(valid) > 0 {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(valid[rand.Intn(len(valid))])
	} else {
		for _, num := range valid {
			fmt.Println(num)
		}
	}
}

func generateRecursive(card, current string, results *[]string) {
	if len(card) == 0 {
		*results = append(*results, current)
		return
	}

	if card[0] == '*' {
		for i := '0'; i <= '9'; i++ {
			generateRecursive(card[1:], current+string(i), results)
		}
	} else {
		generateRecursive(card[1:], current+string(card[0]), results)
	}
}
