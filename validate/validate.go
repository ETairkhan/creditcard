package validate

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Validate(numbers []string) {
	for _, num := range numbers {

		num = strings.ReplaceAll(num, " ", "")

		if len(num) < 13 || len(num) > 19 {
			fmt.Println("INCORRECT")
		} else if IsValidLuhn(num) {
			fmt.Println("OK")
		} else {
			fmt.Println("INCORRECT")
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
		if len(key) == 0 || len(value) == 0 {
			fmt.Printf("Error: Invalid entry in %s file: '%s:%s'\n", fileType, key, value)
			os.Exit(1)
		}
	}
}
