package validators

import (
	"html"
	"regexp"
	"strings"
)

func ValidateDocument(document string) bool {
	document = sanitizeDocument(document)

	switch len(document) {
	case 11:
		return ValidateCPF(document)
	case 14:
		return ValidateCNPJ(document)
	default:
		return false
	}
}

func FormatDocument(document string) string {
	document = sanitizeDocument(document)

	switch len(document) {
	case 11:
		return FormatCPF(document)
	case 14:
		return FormatCNPJ(document)
	default:
		return document
	}
}

func GetDocumentType(document string) string {
	document = sanitizeDocument(document)

	switch len(document) {
	case 11:
		if ValidateCPF(document) {
			return "CPF"
		}
	case 14:
		if ValidateCNPJ(document) {
			return "CNPJ"
		}
	}

	return "inválido"
}

func SanitizeInput(input string) string {
	input = strings.TrimSpace(input)

	input = html.EscapeString(input)

	re := regexp.MustCompile(`[\x00-\x08\x0B\x0C\x0E-\x1F\x7F]`)
	input = re.ReplaceAllString(input, "")

	return input
}

func SanitizeName(name string) string {
	name = SanitizeInput(name)

	re := regexp.MustCompile(`[^a-zA-ZÀ-ÿ\s\-']`)
	name = re.ReplaceAllString(name, "")

	re = regexp.MustCompile(`\s+`)
	name = re.ReplaceAllString(name, " ")

	return strings.TrimSpace(name)
}

func SanitizeEmail(email string) string {
	email = SanitizeInput(email)
	email = strings.ToLower(email)
	email = strings.TrimSpace(email)

	return email
}

func ValidateEmail(email string) bool {
	email = SanitizeEmail(email)

	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
