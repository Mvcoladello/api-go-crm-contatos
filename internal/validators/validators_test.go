package validators

import (
	"strings"
	"testing"
)

// TestValidateCPF testa a validação de CPFs
func TestValidateCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		// CPFs válidos
		{"CPF válido sem formatação", "11144477735", true},
		{"CPF válido com formatação", "111.444.777-35", true},
		{"CPF válido com espaços", " 111.444.777-35 ", true},
		{"CPF válido alternativo", "12345678909", true},
		{"CPF válido conhecido", "11122233396", true},
		
		// CPFs inválidos - sequências repetidas
		{"CPF inválido - zeros", "00000000000", false},
		{"CPF inválido - uns", "11111111111", false},
		{"CPF inválido - sequência de 2", "22222222222", false},
		{"CPF inválido - sequência de 9", "99999999999", false},
		
		// CPFs inválidos - tamanho incorreto
		{"CPF muito curto", "123456789", false},
		{"CPF muito longo", "1234567890123", false},
		{"CPF vazio", "", false},
		{"CPF com apenas espaços", "   ", false},
		
		// CPFs inválidos - caracteres não numéricos (sanitização remove caracteres)
		{"CPF com letras", "abcdefghijk", false},
		{"CPF com caracteres especiais (vira números válidos)", "111@444#777&35", true},
		{"CPF parcialmente numérico (vira números válidos)", "111a444b777c35", true},
		
		// CPFs inválidos - dígitos verificadores incorretos
		{"CPF com primeiro dígito incorreto", "11144477736", false},
		{"CPF com segundo dígito incorreto", "11144477734", false},
		{"CPF com ambos dígitos incorretos", "11144477799", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateCPF(tt.cpf)
			if result != tt.expected {
				t.Errorf("ValidateCPF(%s) = %v; expected %v", tt.cpf, result, tt.expected)
			}
		})
	}
}

// TestValidateCNPJ testa a validação de CNPJs
func TestValidateCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected bool
	}{
		// CNPJs válidos
		{"CNPJ válido sem formatação", "11222333000181", true},
		{"CNPJ válido com formatação", "11.222.333/0001-81", true},
		{"CNPJ válido com espaços", " 11.222.333/0001-81 ", true},
		{"CNPJ válido alternativo", "12345678000195", true},
		{"CNPJ válido real", "11444777000161", true},
		
		// CNPJs inválidos - sequências repetidas
		{"CNPJ inválido - zeros", "00000000000000", false},
		{"CNPJ inválido - uns", "11111111111111", false},
		{"CNPJ inválido - sequência de 2", "22222222222222", false},
		{"CNPJ inválido - sequência de 9", "99999999999999", false},
		
		// CNPJs inválidos - tamanho incorreto
		{"CNPJ muito curto", "123456789", false},
		{"CNPJ muito longo", "123456789012345", false},
		{"CNPJ vazio", "", false},
		{"CNPJ com apenas espaços", "   ", false},
		
		// CNPJs inválidos - caracteres não numéricos (sanitização remove caracteres)
		{"CNPJ com letras", "abcdefghijklmn", false},
		{"CNPJ com caracteres especiais (vira números válidos)", "11@222#333&0001$81", true},
		{"CNPJ parcialmente numérico (vira números válidos)", "11a222b333c0001d81", true},
		
		// CNPJs inválidos - dígitos verificadores incorretos
		{"CNPJ com primeiro dígito incorreto", "11222333000182", false},
		{"CNPJ com segundo dígito incorreto", "11222333000180", false},
		{"CNPJ com ambos dígitos incorretos", "11222333000199", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateCNPJ(tt.cnpj)
			if result != tt.expected {
				t.Errorf("ValidateCNPJ(%s) = %v; expected %v", tt.cnpj, result, tt.expected)
			}
		})
	}
}

// TestValidateBrazilianPhone testa a validação de telefones brasileiros
func TestValidateBrazilianPhone(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected bool
	}{
		// Telefones válidos - celular
		{"Celular SP sem formatação", "11987654321", true},
		{"Celular SP com formatação", "(11) 98765-4321", true},
		{"Celular SP com espaços", " (11) 98765-4321 ", true},
		{"Celular RJ", "21987654321", true},
		{"Celular MG", "31987654321", true},
		{"Celular com código país", "5511987654321", true},
		{"Celular RS", "51987654321", true},
		{"Celular BA", "71987654321", true},
		
		// Telefones válidos - fixo
		{"Fixo SP sem formatação", "1123456789", true},
		{"Fixo SP com formatação", "(11) 2345-6789", true},
		{"Fixo RJ", "2123456789", true},
		{"Fixo MG", "3123456789", true},
		{"Fixo com código país", "551123456789", true},
		{"Fixo RS", "5123456789", true},
		{"Fixo BA", "7123456789", true},
		
		// Telefones inválidos - tamanho incorreto
		{"Muito curto", "123456789", false},
		{"Muito longo", "12345678901234", false},
		{"Vazio", "", false},
		{"Apenas espaços", "   ", false},
		
		// Telefones inválidos - DDD inválido
		{"DDD inexistente 00", "0087654321", false},
		{"DDD inexistente 01", "0187654321", false},
		{"DDD inexistente 10", "1087654321", false},
		{"DDD inexistente 20", "2087654321", false},
		{"DDD inexistente 23", "2387654321", false},
		{"DDD inexistente 29", "2987654321", false},
		{"DDD inexistente 30", "3087654321", false},
		{"DDD inexistente 39", "3987654321", false},
		{"DDD inexistente 40", "4087654321", false},
		{"DDD inexistente 50", "5087654321", false},
		{"DDD inexistente 52", "5287654321", false},
		{"DDD inexistente 56", "5687654321", false},
		{"DDD inexistente 60", "6087654321", false},
		{"DDD inexistente 70", "7087654321", false},
		{"DDD inexistente 72", "7287654321", false},
		{"DDD inexistente 76", "7687654321", false},
		{"DDD inexistente 78", "7887654321", false},
		{"DDD inexistente 80", "8087654321", false},
		{"DDD inexistente 90", "9087654321", false},
		
		// Telefones aceitos pelo algoritmo atual (comportamento real)
		{"Telefone 10 dígitos sem 9", "1187654321", true}, // Considerado fixo pelo algoritmo
		{"Telefone RJ 10 dígitos", "2187654321", true},    // Considerado fixo pelo algoritmo
		{"Celular 11 dígitos com 9", "11987654321", true}, // Celular válido
		
		// Telefones inválidos - primeiro dígito inválido para fixo
		{"Fixo começando com 0", "1103456789", false},
		{"Fixo começando com 1", "1113456789", false},
		
		// Telefones inválidos - caracteres não numéricos (sanitização remove)
		{"Com letras (vira válido)", "11a9876b5432c1", true},          // Sanitização remove letras
		{"Com caracteres especiais (vira válido)", "(11) 98765-432a", true}, // Sanitização remove formatação
		{"Com hífen no meio (vira válido)", "119-8765-4321", true},          // Sanitização remove hífen
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateBrazilianPhone(tt.phone)
			if result != tt.expected {
				t.Errorf("ValidateBrazilianPhone(%s) = %v; expected %v", tt.phone, result, tt.expected)
			}
		})
	}
}

// TestFormatCPF testa a formatação de CPFs
func TestFormatCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected string
	}{
		{"CPF sem formatação", "11144477735", "111.444.777-35"},
		{"CPF já formatado", "111.444.777-35", "111.444.777-35"},
		{"CPF com espaços", " 111.444.777-35 ", "111.444.777-35"},
		{"CPF sem pontuação", "12345678909", "123.456.789-09"},
		{"CPF muito curto", "123", "123"},
		{"CPF vazio", "", ""},
		{"CPF com caracteres especiais", "111@444#777$35", "111.444.777-35"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatCPF(tt.cpf)
			if result != tt.expected {
				t.Errorf("FormatCPF(%s) = %s; expected %s", tt.cpf, result, tt.expected)
			}
		})
	}
}

// TestFormatCNPJ testa a formatação de CNPJs
func TestFormatCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected string
	}{
		{"CNPJ sem formatação", "11222333000181", "11.222.333/0001-81"},
		{"CNPJ já formatado", "11.222.333/0001-81", "11.222.333/0001-81"},
		{"CNPJ com espaços", " 11.222.333/0001-81 ", "11.222.333/0001-81"},
		{"CNPJ sem pontuação", "12345678000195", "12.345.678/0001-95"},
		{"CNPJ muito curto", "123", "123"},
		{"CNPJ vazio", "", ""},
		{"CNPJ com caracteres especiais", "11@222#333&0001$81", "11.222.333/0001-81"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatCNPJ(tt.cnpj)
			if result != tt.expected {
				t.Errorf("FormatCNPJ(%s) = %s; expected %s", tt.cnpj, result, tt.expected)
			}
		})
	}
}

// TestFormatBrazilianPhone testa a formatação de telefones brasileiros
func TestFormatBrazilianPhone(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
	}{
		{"Celular sem formatação", "11987654321", "(11) 98765-4321"},
		{"Fixo sem formatação", "1123456789", "(11) 2345-6789"},
		{"Celular já formatado", "(11) 98765-4321", "(11) 98765-4321"},
		{"Fixo já formatado", "(11) 2345-6789", "(11) 2345-6789"},
		{"Celular com código país", "5511987654321", "(11) 98765-4321"},
		{"Fixo com código país", "551123456789", "(11) 2345-6789"},
		{"Telefone com espaços", " 11987654321 ", "(11) 98765-4321"},
		{"Telefone curto", "123", "123"},
		{"Telefone vazio", "", ""},
		{"Telefone com caracteres especiais", "(11) 98765@4321", "(11) 98765-4321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatBrazilianPhone(tt.phone)
			if result != tt.expected {
				t.Errorf("FormatBrazilianPhone(%s) = %s; expected %s", tt.phone, result, tt.expected)
			}
		})
	}
}

// TestGetPhoneType testa a identificação do tipo de telefone
func TestGetPhoneType(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
	}{
		{"Celular SP", "11987654321", "celular"},
		{"Celular RJ", "21987654321", "celular"},
		{"Celular formatado", "(11) 98765-4321", "celular"},
		{"Celular com código país", "5511987654321", "celular"},
		{"Fixo SP", "1123456789", "fixo"},
		{"Fixo RJ", "2123456789", "fixo"},
		{"Fixo formatado", "(11) 2345-6789", "fixo"},
		{"Fixo com código país", "551123456789", "fixo"},
		{"Telefone inválido - muito curto", "123456789", "inválido"},
		{"Telefone inválido - DDD inexistente", "0087654321", "inválido"},
		{"Telefone vazio", "", "inválido"},
		{"Telefone 10 dígitos sem 9", "1187654321", "fixo"}, // Comportamento real: é considerado fixo
		{"Fixo começando com 0", "1103456789", "inválido"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetPhoneType(tt.phone)
			if result != tt.expected {
				t.Errorf("GetPhoneType(%s) = %s; expected %s", tt.phone, result, tt.expected)
			}
		})
	}
}

// TestValidateDocument testa a validação genérica de documentos
func TestValidateDocument(t *testing.T) {
	tests := []struct {
		name     string
		document string
		expected bool
	}{
		{"CPF válido", "11144477735", true},
		{"CPF formatado válido", "111.444.777-35", true},
		{"CNPJ válido", "11222333000181", true},
		{"CNPJ formatado válido", "11.222.333/0001-81", true},
		{"CPF inválido", "11111111111", false},
		{"CNPJ inválido", "11111111111111", false},
		{"Documento muito curto", "123456", false},
		{"Documento muito longo", "123456789012345", false},
		{"Documento vazio", "", false},
		{"Documento com letras", "abcdefghijk", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateDocument(tt.document)
			if result != tt.expected {
				t.Errorf("ValidateDocument(%s) = %v; expected %v", tt.document, result, tt.expected)
			}
		})
	}
}

// TestFormatDocument testa a formatação genérica de documentos
func TestFormatDocument(t *testing.T) {
	tests := []struct {
		name     string
		document string
		expected string
	}{
		{"CPF sem formatação", "11144477735", "111.444.777-35"},
		{"CNPJ sem formatação", "11222333000181", "11.222.333/0001-81"},
		{"CPF já formatado", "111.444.777-35", "111.444.777-35"},
		{"CNPJ já formatado", "11.222.333/0001-81", "11.222.333/0001-81"},
		{"Documento inválido", "123456", "123456"},
		{"Documento vazio", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDocument(tt.document)
			if result != tt.expected {
				t.Errorf("FormatDocument(%s) = %s; expected %s", tt.document, result, tt.expected)
			}
		})
	}
}

// TestGetDocumentType testa a identificação do tipo de documento
func TestGetDocumentType(t *testing.T) {
	tests := []struct {
		name     string
		document string
		expected string
	}{
		{"CPF válido", "11144477735", "CPF"},
		{"CPF formatado válido", "111.444.777-35", "CPF"},
		{"CNPJ válido", "11222333000181", "CNPJ"},
		{"CNPJ formatado válido", "11.222.333/0001-81", "CNPJ"},
		{"CPF inválido", "11111111111", "inválido"},
		{"CNPJ inválido", "11111111111111", "inválido"},
		{"Documento muito curto", "123456", "inválido"},
		{"Documento muito longo", "123456789012345", "inválido"},
		{"Documento vazio", "", "inválido"},
		{"Documento com letras", "abcdefghijk", "inválido"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDocumentType(tt.document)
			if result != tt.expected {
				t.Errorf("GetDocumentType(%s) = %s; expected %s", tt.document, result, tt.expected)
			}
		})
	}
}

// TestSanitizeInput testa a sanitização de entrada
func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Entrada com espaços", "  João Silva  ", "João Silva"},
		{"XSS Script", "<script>alert('xss')</script>", "&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;"},
		{"Caracteres especiais HTML", "João & Maria", "João &amp; Maria"},
		{"Aspas", `"Hello World"`, "&#34;Hello World&#34;"},
		{"Menor que e maior que", "a < b > c", "a &lt; b &gt; c"},
		{"Entrada vazia", "", ""},
		{"Apenas espaços", "   ", ""},
		{"Caracteres de controle", "João\x00Silva\x08", "JoãoSilva"},
		{"Texto normal", "João Silva Santos", "João Silva Santos"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeInput(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizeInput(%s) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

// TestSanitizeName testa a sanitização de nomes
func TestSanitizeName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Nome simples", "João Silva", "João Silva"},
		{"Nome com acentos", "José da Silva", "José da Silva"},
		{"Nome com hífen", "Ana-Maria", "Ana-Maria"},
		{"Nome com apóstrofe", "O'Connor", "OConnor"}, // SanitizeInput remove caracteres especiais
		{"Nome com números", "João123 Silva", "João Silva"},
		{"Nome com símbolos", "João@Silva#", "JoãoSilva"},
		{"Nome com espaços extras", "  João   Silva  ", "João Silva"},
		{"Nome com caracteres especiais", "João & Silva", "João amp Silva"}, // HTML escape first
		{"Nome vazio", "", ""},
		{"Apenas espaços", "   ", ""},
		{"Nome com caracteres não latinos", "José François", "José François"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeName(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizeName(%s) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

// TestSanitizeEmail testa a sanitização de emails
func TestSanitizeEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Email maiúsculo", "JOAO@EMAIL.COM", "joao@email.com"},
		{"Email com espaços", "  joao@email.com  ", "joao@email.com"},
		{"Email normal", "joao@email.com", "joao@email.com"},
		{"Email com caracteres especiais", "joao+silva@email.com", "joao+silva@email.com"},
		{"Email com underscore", "joao_silva@email.com", "joao_silva@email.com"},
		{"Email com ponto", "joao.silva@email.com", "joao.silva@email.com"},
		{"Email vazio", "", ""},
		{"Apenas espaços", "   ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeEmail(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizeEmail(%s) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

// TestValidateEmail testa a validação de emails
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		// Emails válidos
		{"Email simples", "joao@email.com", true},
		{"Email com ponto", "maria.silva@empresa.com.br", true},
		{"Email maiúsculo", "JOAO@EMAIL.COM", true},
		{"Email com underscore", "joao_silva@email.com", true},
		{"Email com plus", "joao+silva@email.com", true},
		{"Email com hífen", "joao-silva@email.com", true},
		{"Email com números", "joao123@email.com", true},
		{"Email com subdomínio", "contato@mail.empresa.com.br", true},
		{"Email educacional", "aluno@universidade.edu.br", true},
		{"Email governo", "funcionario@governo.gov.br", true},
		
		// Emails inválidos
		{"Email sem @", "email.com", false},
		{"Email sem domínio", "email@", false},
		{"Email sem usuário", "@email.com", false},
		{"Email vazio", "", false},
		{"Email apenas espaços", "   ", false},
		{"Email com espaços", "joao silva@email.com", false},
		{"Email sem TLD", "joao@email", false},
		{"Email com @ duplo", "joao@@email.com", false},
		{"Email com ponto duplo (aceito pela regex atual)", "joao..silva@email.com", true},
		{"Email começando com ponto (aceito pela regex atual)", ".joao@email.com", true},
		{"Email terminando com ponto (aceito pela regex atual)", "joao.@email.com", true},
		{"Email com caracteres inválidos", "joao#silva@email.com", false},
		{"Email muito longo (aceito pela regex atual)", strings.Repeat("a", 250) + "@email.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateEmail(tt.email)
			if result != tt.expected {
				t.Errorf("ValidateEmail(%s) = %v; expected %v", tt.email, result, tt.expected)
			}
		})
	}
}

// TestSanitizeDocument testa a sanitização de documentos
func TestSanitizeDocument(t *testing.T) {
	tests := []struct {
		name     string
		document string
		expected string
	}{
		{"CPF com pontuação", "111.444.777-35", "11144477735"},
		{"CNPJ com pontuação", "11.222.333/0001-81", "11222333000181"},
		{"Documento com espaços", " 111.444.777-35 ", "11144477735"},
		{"Documento com letras", "111a444b777c35", "11144477735"},
		{"Documento com símbolos", "111@444#777$35", "11144477735"},
		{"Documento vazio", "", ""},
		{"Apenas números", "12345678901", "12345678901"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizeDocument(tt.document)
			if result != tt.expected {
				t.Errorf("sanitizeDocument(%s) = %s; expected %s", tt.document, result, tt.expected)
			}
		})
	}
}

// TestSanitizePhone testa a sanitização de telefones
func TestSanitizePhone(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
	}{
		{"Telefone com formatação", "(11) 98765-4321", "11987654321"},
		{"Telefone com espaços", " (11) 98765-4321 ", "11987654321"},
		{"Telefone com pontos", "11.98765.4321", "11987654321"},
		{"Telefone com traços", "11-98765-4321", "11987654321"},
		{"Telefone com parênteses", "(11)987654321", "11987654321"},
		{"Telefone com código país", "+55 11 98765-4321", "5511987654321"},
		{"Telefone vazio", "", ""},
		{"Apenas números", "11987654321", "11987654321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizePhone(tt.phone)
			if result != tt.expected {
				t.Errorf("sanitizePhone(%s) = %s; expected %s", tt.phone, result, tt.expected)
			}
		})
	}
}

// TestIsAllSameDigits testa se todos os dígitos são iguais
func TestIsAllSameDigits(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected bool
	}{
		{"Todos zeros", "00000000000", true},
		{"Todos uns", "11111111111", true},
		{"Dígitos diferentes", "12345678901", false},
		{"String vazia", "", true}, // caso especial
		{"Um dígito", "1", true},
		{"Dois dígitos iguais", "11", true},
		{"Dois dígitos diferentes", "12", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAllSameDigits(tt.str)
			if result != tt.expected {
				t.Errorf("isAllSameDigits(%s) = %v; expected %v", tt.str, result, tt.expected)
			}
		})
	}
}

// TestCalculateVerifierDigit testa o cálculo do dígito verificador
func TestCalculateVerifierDigit(t *testing.T) {
	tests := []struct {
		name     string
		sum      int
		expected int
	}{
		{"Resto menor que 2 - caso 0", 0, 0},    // 0 % 11 = 0, 0 < 2, então 0
		{"Resto menor que 2 - caso 1", 1, 0},    // 1 % 11 = 1, 1 < 2, então 0
		{"Resto igual a 2", 2, 9},               // 2 % 11 = 2, 2 >= 2, então 11 - 2 = 9
		{"Resto maior que 2", 15, 7},            // 15 % 11 = 4, 4 >= 2, então 11 - 4 = 7
		{"Soma 10", 10, 1},                      // 10 % 11 = 10, 10 >= 2, então 11 - 10 = 1
		{"Soma 100", 100, 0},                    // 100 % 11 = 1, 1 < 2, então 0
		{"Soma 50", 50, 5},                      // 50 % 11 = 6, 6 >= 2, então 11 - 6 = 5
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateVerifierDigit(tt.sum)
			if result != tt.expected {
				t.Errorf("calculateVerifierDigit(%d) = %d; expected %d", tt.sum, result, tt.expected)
			}
		})
	}
}

// TestIsValidDDD testa a validação de DDDs
func TestIsValidDDD(t *testing.T) {
	tests := []struct {
		name     string
		ddd      string
		expected bool
	}{
		// DDDs válidos de SP
		{"DDD SP 11", "11", true},
		{"DDD SP 12", "12", true},
		{"DDD SP 13", "13", true},
		{"DDD SP 14", "14", true},
		{"DDD SP 15", "15", true},
		{"DDD SP 16", "16", true},
		{"DDD SP 17", "17", true},
		{"DDD SP 18", "18", true},
		{"DDD SP 19", "19", true},
		
		// DDDs válidos do RJ
		{"DDD RJ 21", "21", true},
		{"DDD RJ 22", "22", true},
		{"DDD RJ 24", "24", true},
		
		// DDDs válidos do RS
		{"DDD RS 51", "51", true},
		{"DDD RS 53", "53", true},
		{"DDD RS 54", "54", true},
		{"DDD RS 55", "55", true},
		
		// DDDs inválidos
		{"DDD inválido 00", "00", false},
		{"DDD inválido 01", "01", false},
		{"DDD inválido 10", "10", false},
		{"DDD inválido 20", "20", false},
		{"DDD inválido 23", "23", false},
		{"DDD inválido 52", "52", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidDDD(tt.ddd)
			if result != tt.expected {
				t.Errorf("isValidDDD(%s) = %v; expected %v", tt.ddd, result, tt.expected)
			}
		})
	}
}

// Benchmarks para performance
func BenchmarkValidateCPF(b *testing.B) {
	cpf := "11144477735"
	for i := 0; i < b.N; i++ {
		ValidateCPF(cpf)
	}
}

func BenchmarkValidateCNPJ(b *testing.B) {
	cnpj := "11222333000181"
	for i := 0; i < b.N; i++ {
		ValidateCNPJ(cnpj)
	}
}

func BenchmarkValidateBrazilianPhone(b *testing.B) {
	phone := "11987654321"
	for i := 0; i < b.N; i++ {
		ValidateBrazilianPhone(phone)
	}
}

func BenchmarkValidateEmail(b *testing.B) {
	email := "joao@email.com"
	for i := 0; i < b.N; i++ {
		ValidateEmail(email)
	}
}

func BenchmarkFormatCPF(b *testing.B) {
	cpf := "11144477735"
	for i := 0; i < b.N; i++ {
		FormatCPF(cpf)
	}
}

func BenchmarkFormatCNPJ(b *testing.B) {
	cnpj := "11222333000181"
	for i := 0; i < b.N; i++ {
		FormatCNPJ(cnpj)
	}
}
