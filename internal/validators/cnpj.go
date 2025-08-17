package validators

import (
	"strconv"
)

func ValidateCNPJ(cnpj string) bool {
	cnpj = sanitizeDocument(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	if isAllSameDigits(cnpj) {
		return false
	}

	digits := make([]int, 14)
	for i, char := range cnpj {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digits[i] = digit
	}

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i := 0; i < 12; i++ {
		sum += digits[i] * weights1[i]
	}
	firstDigit := calculateVerifierDigit(sum)

	if digits[12] != firstDigit {
		return false
	}

	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sum = 0
	for i := 0; i < 13; i++ {
		sum += digits[i] * weights2[i]
	}
	secondDigit := calculateVerifierDigit(sum)

	return digits[13] == secondDigit
}

func FormatCNPJ(cnpj string) string {
	cnpj = sanitizeDocument(cnpj)
	if len(cnpj) != 14 {
		return cnpj
	}
	return cnpj[:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:]
}
