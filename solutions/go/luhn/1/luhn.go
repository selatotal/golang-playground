package luhn

import "unicode"

func Valid(id string) bool {
	sum := 0
	double := false
	digits := 0

	for i := len(id) - 1; i >= 0; i-- {
		char := rune(id[i])

		if char == ' ' {
			continue
		}

		if !unicode.IsDigit(char) {
			return false
		}

		digit := int(char - '0')
		digits++

		if double {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return digits > 1 && sum%10 == 0
}