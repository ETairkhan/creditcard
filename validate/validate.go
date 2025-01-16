package validate

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Validate(numbers []string) {
	if numbers[0] == "" {
		os.Exit(1)
	}
	for _, num := range numbers {

		num = strings.ReplaceAll(num, " ", "")

		if len(num) < 13 || len(num) > 19 {
			fmt.Println("INCORRECT")
			os.Exit(1)
		} else if IsValidLuhn(num) {
			fmt.Println("OK")
		} else {
			fmt.Println("INCORRECT")
			os.Exit(1)
		}
	}
}

func IsValidLuhn(card string) bool {
	sum, double := 0, false
	for i := len(card) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(card[i]))
		if err != nil {
			return false
		}
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	if sum == 0 {
		return false
	}
	return sum%10 == 0
}

func ValidateData(data map[string]string, fileType string) {
	for key, value := range data {
		if key == "" || value == "" {
			fmt.Printf("Error: Invalid entry in %s file: '%s:%s'\n", fileType, key, value)
			os.Exit(1)
		}
	}
}
