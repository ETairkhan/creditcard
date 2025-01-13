package validate

import (
	"fmt"
	"strconv"
)

func Validate(numbers []string) {
	for _, num := range numbers {
		if len(num) < 13{
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
