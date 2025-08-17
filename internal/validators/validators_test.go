package validators

import (
	"testing"
)

func TestValidateCPF(t *testing.T) {
	tests := []struct {
		cpf      string
		expected bool
	}{
		{"11144477735", true},
		{"111.444.777-35", true},
		{"00000000000", false},
		{"11111111111", false},
		{"123456789", false},
		{"1234567890123", false},
		{"abcdefghijk", false},
	}

	for _, test := range tests {
		result := ValidateCPF(test.cpf)
		if result != test.expected {
			t.Errorf("ValidateCPF(%s) = %v; expected %v", test.cpf, result, test.expected)
		}
	}
}

func TestValidateCNPJ(t *testing.T) {
	tests := []struct {
		cnpj     string
		expected bool
	}{
		{"11222333000181", true},
		{"11.222.333/0001-81", true},
		{"00000000000000", false},
		{"11111111111111", false},
		{"123456789", false},
		{"abcdefghijklmn", false},
	}

	for _, test := range tests {
		result := ValidateCNPJ(test.cnpj)
		if result != test.expected {
			t.Errorf("ValidateCNPJ(%s) = %v; expected %v", test.cnpj, result, test.expected)
		}
	}
}

func TestValidateBrazilianPhone(t *testing.T) {
	tests := []struct {
		phone    string
		expected bool
	}{
		{"11987654321", true},
		{"1123456789", true},
		{"(11) 98765-4321", true},
		{"(11) 2345-6789", true},
		{"5511987654321", true},
		{"123456789", false},
		{"12345678901234", false},
		{"0187654321", false},
		{"1109876543", false},
		{"1119876543", false},
	}

	for _, test := range tests {
		result := ValidateBrazilianPhone(test.phone)
		if result != test.expected {
			t.Errorf("ValidateBrazilianPhone(%s) = %v; expected %v", test.phone, result, test.expected)
		}
	}
}

func TestFormatCPF(t *testing.T) {
	tests := []struct {
		cpf      string
		expected string
	}{
		{"11144477735", "111.444.777-35"},
		{"111.444.777-35", "111.444.777-35"},
		{"123", "123"},
	}

	for _, test := range tests {
		result := FormatCPF(test.cpf)
		if result != test.expected {
			t.Errorf("FormatCPF(%s) = %s; expected %s", test.cpf, result, test.expected)
		}
	}
}

func TestFormatBrazilianPhone(t *testing.T) {
	tests := []struct {
		phone    string
		expected string
	}{
		{"11987654321", "(11) 98765-4321"},
		{"1123456789", "(11) 2345-6789"},
		{"5511987654321", "(11) 98765-4321"},
		{"123", "123"},
	}

	for _, test := range tests {
		result := FormatBrazilianPhone(test.phone)
		if result != test.expected {
			t.Errorf("FormatBrazilianPhone(%s) = %s; expected %s", test.phone, result, test.expected)
		}
	}
}

func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  João Silva  ", "João Silva"},
		{"<script>alert('xss')</script>", "&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;"},
		{"João & Maria", "João &amp; Maria"},
	}

	for _, test := range tests {
		result := SanitizeInput(test.input)
		if result != test.expected {
			t.Errorf("SanitizeInput(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"joao@email.com", true},
		{"maria.silva@empresa.com.br", true},
		{"JOAO@EMAIL.COM", true},
		{"email@", false},
		{"@email.com", false},
		{"email.com", false},
		{"", false},
	}

	for _, test := range tests {
		result := ValidateEmail(test.email)
		if result != test.expected {
			t.Errorf("ValidateEmail(%s) = %v; expected %v", test.email, result, test.expected)
		}
	}
}
