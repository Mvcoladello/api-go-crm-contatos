// Package validators - Exemplos de uso
//
// Este arquivo demonstra como usar os validadores criados
package validators

import "fmt"

// ExampleUsage demonstra o uso dos validadores
func ExampleUsage() {
	// Exemplo de validação de CPF
	cpf := "111.444.777-35"
	if ValidateCPF(cpf) {
		fmt.Printf("CPF %s é válido\n", cpf)
		fmt.Printf("CPF formatado: %s\n", FormatCPF(cpf))
	}

	// Exemplo de validação de CNPJ
	cnpj := "11.222.333/0001-81"
	if ValidateCNPJ(cnpj) {
		fmt.Printf("CNPJ %s é válido\n", cnpj)
		fmt.Printf("CNPJ formatado: %s\n", FormatCNPJ(cnpj))
	}

	// Exemplo de validação de telefone
	telefone := "11987654321"
	if ValidateBrazilianPhone(telefone) {
		fmt.Printf("Telefone %s é válido\n", telefone)
		fmt.Printf("Telefone formatado: %s\n", FormatBrazilianPhone(telefone))
		fmt.Printf("Tipo do telefone: %s\n", GetPhoneType(telefone))
	}

	// Exemplo de validação de documento genérico
	documento := "12345678901"
	if ValidateDocument(documento) {
		fmt.Printf("Documento %s é válido\n", documento)
		fmt.Printf("Tipo: %s\n", GetDocumentType(documento))
		fmt.Printf("Formatado: %s\n", FormatDocument(documento))
	}

	// Exemplo de sanitização
	entrada := "  João <script>alert('xss')</script> Silva  "
	saida := SanitizeInput(entrada)
	fmt.Printf("Entrada: %s\n", entrada)
	fmt.Printf("Saída sanitizada: %s\n", saida)

	// Exemplo de sanitização de nome
	nome := "João123 $$$ Silva"
	nomeLimpo := SanitizeName(nome)
	fmt.Printf("Nome original: %s\n", nome)
	fmt.Printf("Nome sanitizado: %s\n", nomeLimpo)

	// Exemplo de validação de email
	email := "JOAO@EMAIL.COM"
	if ValidateEmail(email) {
		emailLimpo := SanitizeEmail(email)
		fmt.Printf("Email válido: %s\n", emailLimpo)
	}
}
