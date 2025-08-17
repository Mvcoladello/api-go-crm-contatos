# API CRM Contatos

API REST para gerenciamento de contatos com validação de CPF/CNPJ e telefone brasileiro.

## Funcionalidades

- ✅ Listar todos os contatos
- ✅ Buscar contato por ID
- ✅ Criar novo contato
- ✅ Deletar contato
- ✅ Validação de CPF/CNPJ
- ✅ Validação de telefone brasileiro
- ✅ Validação de email

## Endpoints

### 1. Listar contatos
```
GET /api/v1/contatos
```

**Resposta de sucesso (200):**
```json
{
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "nome": "João Silva",
      "email": "joao@example.com",
      "cpf_cnpj": "123.456.789-00",
      "telefone": "(11) 99999-9999",
      "created_at": "2025-08-17T10:00:00Z",
      "updated_at": "2025-08-17T10:00:00Z"
    }
  ],
  "total": 1
}
```

### 2. Buscar contato por ID
```
GET /api/v1/contatos/:id
```

**Parâmetros:**
- `id` (UUID): ID do contato

**Resposta de sucesso (200):**
```json
{
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "nome": "João Silva",
    "email": "joao@example.com",
    "cpf_cnpj": "123.456.789-00",
    "telefone": "(11) 99999-9999",
    "created_at": "2025-08-17T10:00:00Z",
    "updated_at": "2025-08-17T10:00:00Z"
  }
}
```

**Resposta de erro (404):**
```json
{
  "error": "Contato não encontrado"
}
```

### 3. Criar contato
```
POST /api/v1/contatos
```

**Corpo da requisição:**
```json
{
  "nome": "João Silva",
  "email": "joao@example.com",
  "cpf_cnpj": "12345678900",
  "telefone": "11999999999"
}
```

**Resposta de sucesso (201):**
```json
{
  "message": "Contato criado com sucesso",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "nome": "João Silva",
    "email": "joao@example.com",
    "cpf_cnpj": "123.456.789-00",
    "telefone": "(11) 99999-9999",
    "created_at": "2025-08-17T10:00:00Z",
    "updated_at": "2025-08-17T10:00:00Z"
  }
}
```

**Resposta de erro (400):**
```json
{
  "error": "Dados inválidos",
  "details": "CPF/CNPJ inválido"
}
```

**Resposta de erro (409):**
```json
{
  "error": "Email ou CPF/CNPJ já cadastrado"
}
```

### 4. Deletar contato
```
DELETE /api/v1/contatos/:id
```

**Parâmetros:**
- `id` (UUID): ID do contato

**Resposta de sucesso (200):**
```json
{
  "message": "Contato deletado com sucesso"
}
```

**Resposta de erro (404):**
```json
{
  "error": "Contato não encontrado"
}
```

## Validações

### CPF/CNPJ
- Aceita CPF (11 dígitos) ou CNPJ (14 dígitos)
- Validação matemática dos dígitos verificadores
- Formatação automática com pontos, traços e barras

### Telefone
- Aceita telefones brasileiros fixos e celulares
- Formatos aceitos: DDD + número (8 ou 9 dígitos)
- Formatação automática com parênteses e traços

### Email
- Validação de formato de email padrão
- Campo único no banco de dados

## Como executar

1. **Instalar dependências:**
```bash
go mod download
```

2. **Executar a aplicação:**
```bash
go run main.go
```

3. **A API estará disponível em:**
```
http://localhost:3000
```

4. **Verificar saúde da API:**
```
GET http://localhost:3000/health
```

## Banco de Dados

A aplicação usa SQLite como banco de dados padrão, criando automaticamente o arquivo `crm_contatos.db` na raiz do projeto.

## Estrutura do Projeto

```
.
├── main.go
├── go.mod
├── go.sum
└── internal/
    ├── handlers/
    │   ├── contact_handler.go
    │   └── routes.go
    ├── models/
    │   └── contact.go
    ├── services/
    │   └── contact_service.go
    └── validators/
        ├── cpf.go
        ├── cnpj.go
        ├── phone.go
        └── sanitizer.go
```
