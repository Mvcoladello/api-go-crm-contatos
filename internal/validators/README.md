# Validadores para API Go CRM Contatos

Este pacote contém validadores específicos para dados brasileiros, incluindo CPF, CNPJ e telefones, além de funcionalidades de sanitização para prevenir XSS e garantir a integridade dos dados.

## Funcionalidades

### Validação de CPF
- ✅ Validação usando algoritmo oficial do CPF
- ✅ Formatação automática (XXX.XXX.XXX-XX)
- ✅ Sanitização de entrada (remove caracteres não numéricos)
- ✅ Verificação de CPFs inválidos conhecidos (todos os dígitos iguais)

### Validação de CNPJ
- ✅ Validação usando algoritmo oficial do CNPJ
- ✅ Formatação automática (XX.XXX.XXX/XXXX-XX)
- ✅ Sanitização de entrada (remove caracteres não numéricos)
- ✅ Verificação de CNPJs inválidos conhecidos

### Validação de Telefone Brasileiro
- ✅ Suporte a telefones fixos (10 dígitos) e celulares (11 dígitos)
- ✅ Validação de DDDs válidos no Brasil
- ✅ Formatação automática ((XX) XXXXX-XXXX ou (XX) XXXX-XXXX)
- ✅ Suporte a código do país (+55)
- ✅ Identificação do tipo (celular/fixo)

### Sanitização e Segurança
- ✅ Sanitização de entrada para prevenir XSS
- ✅ Escape de caracteres HTML
- ✅ Limpeza de caracteres de controle
- ✅ Sanitização específica para nomes e emails

## Uso

### Validação de CPF
```go
import "github.com/mvcoladello/api-go-crm-contatos/internal/validators"

// Validar CPF
if validators.ValidateCPF("111.444.777-35") {
    fmt.Println("CPF válido!")
}

// Formatar CPF
cpfFormatado := validators.FormatCPF("11144477735")
// Resultado: "111.444.777-35"
```

### Validação de CNPJ
```go
// Validar CNPJ
if validators.ValidateCNPJ("11.222.333/0001-81") {
    fmt.Println("CNPJ válido!")
}

// Formatar CNPJ
cnpjFormatado := validators.FormatCNPJ("11222333000181")
// Resultado: "11.222.333/0001-81"
```

### Validação de Telefone
```go
// Validar telefone
if validators.ValidateBrazilianPhone("11987654321") {
    fmt.Println("Telefone válido!")
}

// Formatar telefone
telefoneFormatado := validators.FormatBrazilianPhone("11987654321")
// Resultado: "(11) 98765-4321"

// Identificar tipo
tipo := validators.GetPhoneType("11987654321")
// Resultado: "celular"
```

### Validação Genérica de Documento
```go
// Validar automaticamente CPF ou CNPJ
if validators.ValidateDocument("11144477735") {
    tipo := validators.GetDocumentType("11144477735")
    fmt.Printf("Documento válido do tipo: %s\n", tipo)
}
```

### Sanitização
```go
// Sanitizar entrada geral
entrada := "  João <script>alert('xss')</script>  "
saida := validators.SanitizeInput(entrada)
// Resultado: "João &lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;"

// Sanitizar nome
nome := validators.SanitizeName("João123 Silva")
// Resultado: "João Silva"

// Sanitizar email
email := validators.SanitizeEmail("  JOAO@EMAIL.COM  ")
// Resultado: "joao@email.com"
```

## Integração com Model Contact

O model `Contact` já está integrado com os validadores e automaticamente:
- Sanitiza todos os campos de entrada
- Valida CPF/CNPJ usando os algoritmos oficiais
- Valida telefones brasileiros
- Formata automaticamente os dados após validação
- Retorna erros descritivos em caso de dados inválidos

## Executar Testes

```bash
go test ./internal/validators/ -v
```

## Arquivos

- `cpf.go` - Validação e formatação de CPF
- `cnpj.go` - Validação e formatação de CNPJ
- `phone.go` - Validação e formatação de telefones brasileiros
- `sanitizer.go` - Funções de sanitização e validação geral
- `validators_test.go` - Testes unitários
- `examples.go` - Exemplos de uso

## Segurança

- ✅ Prevenção contra XSS através de escape de HTML
- ✅ Sanitização de caracteres de controle
- ✅ Validação rigorosa de formatos
- ✅ Limpeza automática de dados de entrada
