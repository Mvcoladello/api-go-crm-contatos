# Respostas Padronizadas da API

Este documento descreve o formato padronizado das respostas da API CRM Contatos.

## Formato de Resposta de Sucesso

### Estrutura Base
```json
{
  "message": "string (opcional)",
  "data": "object|array (opcional)",
  "total": "number (opcional, para listas)",
  "timestamp": "string (ISO 8601)",
  "meta": "object (opcional)"
}
```

### Exemplos de Sucesso

**Listagem de contatos:**
```json
{
  "data": [
    {
      "id": "19dd67af-495c-425a-bc01-dc03a6f3606f",
      "nome": "Ana Santos",
      "email": "ana@example.com",
      "cpf_cnpj": "123.456.789-09",
      "telefone": "(11) 98765-4321",
      "created_at": "2025-08-17T01:18:29.23753-04:00",
      "updated_at": "2025-08-17T01:18:29.23753-04:00",
      "deleted_at": null
    }
  ],
  "total": 1,
  "timestamp": "2025-08-17T01:18:29.238172-04:00"
}
```

**Criação de contato:**
```json
{
  "message": "Contato criado com sucesso",
  "data": {
    "id": "19dd67af-495c-425a-bc01-dc03a6f3606f",
    "nome": "Ana Santos",
    "email": "ana@example.com",
    "cpf_cnpj": "123.456.789-09",
    "telefone": "(11) 98765-4321",
    "created_at": "2025-08-17T01:18:29.23753-04:00",
    "updated_at": "2025-08-17T01:18:29.23753-04:00",
    "deleted_at": null
  },
  "timestamp": "2025-08-17T01:18:29.238172-04:00"
}
```

## Formato de Resposta de Erro

### Estrutura Base
```json
{
  "error": "string",
  "code": "string",
  "details": "string (opcional)",
  "timestamp": "string (ISO 8601)",
  "path": "string",
  "method": "string",
  "data": "object (opcional)"
}
```

### Códigos de Erro

| Código | Descrição | Status HTTP |
|--------|-----------|-------------|
| `VALIDATION_ERROR` | Erro de validação de dados | 400 |
| `BAD_REQUEST` | Requisição malformada | 400 |
| `NOT_FOUND` | Recurso não encontrado | 404 |
| `CONFLICT` | Conflito de dados (duplicação) | 409 |
| `INTERNAL_SERVER_ERROR` | Erro interno do servidor | 500 |
| `UNAUTHORIZED` | Não autorizado | 401 |
| `FORBIDDEN` | Acesso negado | 403 |
| `INVALID_FORMAT` | Formato de dados inválido | 400 |
| `DUPLICATE_RESOURCE` | Recurso duplicado | 409 |

### Exemplos de Erro

**Erro de validação:**
```json
{
  "error": "Nome é obrigatório",
  "code": "VALIDATION_ERROR",
  "timestamp": "2025-08-17T01:18:16.749205-04:00",
  "path": "/api/v1/contatos",
  "method": "POST"
}
```

**Erro de ID inválido:**
```json
{
  "error": "ID deve ser um UUID válido",
  "code": "BAD_REQUEST",
  "timestamp": "2025-08-17T01:18:40.963865-04:00",
  "path": "/api/v1/contatos/invalid-id",
  "method": "GET"
}
```

**Erro de duplicação:**
```json
{
  "error": "Email ou CPF/CNPJ já cadastrado",
  "code": "CONFLICT",
  "timestamp": "2025-08-17T01:18:38.733762-04:00",
  "path": "/api/v1/contatos",
  "method": "POST"
}
```

**Erro interno do servidor:**
```json
{
  "error": "Erro interno do servidor",
  "code": "INTERNAL_SERVER_ERROR",
  "details": "UNIQUE constraint failed: contacts.email",
  "timestamp": "2025-08-17T01:18:49.227395-04:00",
  "path": "/api/v1/contatos",
  "method": "POST"
}
```

## Vantagens da Padronização

### 1. **Consistência**
- Todas as respostas seguem o mesmo formato
- Facilita o desenvolvimento do frontend
- Reduz ambiguidade na documentação

### 2. **Rastreabilidade**
- Timestamp em todas as respostas
- Path e method para identificar origem do erro
- Códigos estruturados para diferentes tipos de erro

### 3. **Debugging**
- Campo `details` para informações técnicas
- Logs automáticos com contexto completo
- Identificação clara do tipo de erro

### 4. **Experiência do Desenvolvedor**
- Respostas previsíveis e estruturadas
- Códigos de erro semânticos
- Informações suficientes para tratamento de erro

## Implementação

### Handlers
Os handlers utilizam funções utilitárias para garantir consistência:

```go
// Sucesso
return utils.SendSuccessResponse(c, http.StatusOK, "Mensagem", dados)

// Erro de validação
return utils.SendValidationError(c, "Mensagem de erro")

// Erro de conflito
return utils.SendConflictError(c, "Recurso já existe")
```

### Middlewares
Os middlewares também seguem o padrão de respostas:

```go
// Erro de validação no middleware
return utils.SendValidationError(c, "Campo obrigatório")

// Erro de formato
return utils.SendBadRequestError(c, "Formato inválido")
```

Esta padronização garante que toda a API tenha respostas consistentes e facilita tanto o desenvolvimento quanto o debugging.
