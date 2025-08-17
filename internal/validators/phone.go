package validators

import (
	"regexp"
	"strings"
)

func ValidateBrazilianPhone(phone string) bool {
	phone = sanitizePhone(phone)

	phone = strings.TrimPrefix(phone, "55")

	if len(phone) != 10 && len(phone) != 11 {
		return false
	}

	if len(phone) >= 2 {
		ddd := phone[:2]
		if !isValidDDD(ddd) {
			return false
		}
	}

	if len(phone) == 11 {
		if phone[2] != '9' {
			return false
		}
	}

	if len(phone) == 10 {
		if phone[2] == '0' || phone[2] == '1' {
			return false
		}
	}

	return true
}

func FormatBrazilianPhone(phone string) string {
	phone = sanitizePhone(phone)

	phone = strings.TrimPrefix(phone, "55")

	if len(phone) == 11 {
		return "(" + phone[:2] + ") " + phone[2:7] + "-" + phone[7:]
	} else if len(phone) == 10 {
		return "(" + phone[:2] + ") " + phone[2:6] + "-" + phone[6:]
	}

	return phone
}

func GetPhoneType(phone string) string {
	phone = sanitizePhone(phone)

	phone = strings.TrimPrefix(phone, "55")

	if !ValidateBrazilianPhone(phone) {
		return "inválido"
	}

	if len(phone) == 11 && phone[2] == '9' {
		return "celular"
	} else if len(phone) == 10 {
		return "fixo"
	}

	return "inválido"
}

func sanitizePhone(phone string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(phone, "")
}

func isValidDDD(ddd string) bool {
	validDDDs := map[string]bool{
		"11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true,
		"21": true, "22": true, "24": true, "27": true, "28": true,
		"31": true, "32": true, "33": true, "34": true, "35": true, "37": true, "38": true,
		"41": true, "42": true, "43": true, "44": true, "45": true, "46": true,
		"47": true, "48": true, "49": true,
		"51": true, "53": true, "54": true, "55": true,
		"61": true, "62": true, "63": true, "64": true, "65": true, "66": true, "67": true, "68": true, "69": true,
		"71": true, "73": true, "74": true, "75": true, "77": true, "79": true,
		"81": true, "82": true, "83": true, "84": true, "85": true, "86": true, "87": true, "88": true, "89": true,
		"91": true, "92": true, "93": true, "94": true, "95": true, "96": true, "97": true, "98": true, "99": true,
	}

	return validDDDs[ddd]
}
