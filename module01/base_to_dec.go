package module01

import "strings"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	answer := 0
	value = strings.ToUpper(value)

	for i := 0; i < len(value); i++ {
		ch := value[i]
		var digit int

		if ch >= '0' && ch <= '9' {
			digit = int(ch - '0')
		} else if ch >= 'A' && ch <= 'Z' {
			digit = int(ch - 'A' + 10)
		}

		// Multiply current total by base and add digit
		answer = answer*base + digit
	}

	return answer
}
