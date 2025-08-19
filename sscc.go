package utility

import "fmt"

// calculate a check digit for sscc
// https://www.gs1.org/services/how-calculate-check-digit-manually
// Codes are padded with zeros up to 17 positions or truncated to the first 17 positions
func Sscc(code string) (out string, err error) {
	// Validate that code contains only digits
	if code == "" {
		return "", fmt.Errorf("code is empty string")
	}
	for i, ch := range code {
		if ch < '0' || ch > '9' {
			return "", fmt.Errorf("invalid character '%c' at position %d", ch, i)
		}
	}
	switch {
	case len(code) > 17:
		code = code[:17]
	case len(code) < 17:
		code = fmt.Sprintf("%017s", code)
	}
	sum := 0
	for i := range code {
		n := code[i] - '0'
		if i%2 == 0 {
			n *= 3
			sum += int(n)
		} else {
			sum += int(n)
		}
	}
	return fmt.Sprintf("00%s%d", code, roundUp(sum)-sum), nil
}

func roundUp(val int) int {
	return 10 * ((val + 9) / 10)
}
