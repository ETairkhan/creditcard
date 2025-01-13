package validate

import (
	"errors"
	"fmt"
	"strconv"
)

func Validate(numbers []string) error {
	for _, num := range numbers {
		if len(num) < 13 {
			return errors.New("INCORRECT: number too short")
		}
		if IsValidLuhn(num) {
			fmt.Println("OK")
		} else {
			return errors.New("INCORRECT: invalid number")
		}
	}
	return nil
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
