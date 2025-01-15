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
	if strings.Count(card, "*") > 4 {
		fmt.Println("1")
		fmt.Println("$ echo $? \n1")
		os.Exit(1)
	}

	results := []string{}
	generateRecursive(card, "", &results)

	valid := []string{}
	for _, num := range results {
		if validate.IsValidLuhn(num) {
			valid = append(valid, num)
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
