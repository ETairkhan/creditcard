package validate

import (
	"fmt"
	"os"
	"strconv"
)

func Validate(numbers []string) {
	for _, num := range numbers {
		if len(num) < 13 {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
		if isValidLuhn(num) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
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
	return sum%10 == 0
}
