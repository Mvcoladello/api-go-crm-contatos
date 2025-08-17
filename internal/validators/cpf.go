package validators

import (
	"regexp"
	"strconv"
)

func ValidateCPF(cpf string) bool {
	cpf = sanitizeDocument(cpf)

	if len(cpf) != 11 {
		return false
	}

	if isAllSameDigits(cpf) {
		return false
	}

	digits := make([]int, 11)
	for i, char := range cpf {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digits[i] = digit
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	firstDigit := calculateVerifierDigit(sum)

	if digits[9] != firstDigit {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (11 - i)
	}
	secondDigit := calculateVerifierDigit(sum)

	return digits[10] == secondDigit
}

func FormatCPF(cpf string) string {
	cpf = sanitizeDocument(cpf)
	if len(cpf) != 11 {
		return cpf
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}

func sanitizeDocument(document string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(document, "")
}

func isAllSameDigits(str string) bool {
	for i := 1; i < len(str); i++ {
		if str[i] != str[0] {
			return false
		}
	}
	return true
}

func calculateVerifierDigit(sum int) int {
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
