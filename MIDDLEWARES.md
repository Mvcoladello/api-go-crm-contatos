# Middlewares

Este documento descreve os middlewares implementados na API CRM Contatos.

## LoggingMiddleware

Middleware responsável por registrar logs detalhados de todas as requisições HTTP.

**Funcionalidades:**
- Registra início e fim de cada requisição
- Inclui método HTTP, path, IP do cliente
- Mede e registra tempo de execução
- Registra código de status HTTP da resposta

**Logs gerados:**
```
[POST] /api/v1/contatos 127.0.0.1 - Iniciado
[POST] /api/v1/contatos 127.0.0.1 - Finalizado em 3.366542ms - Status: 201
```

## TracingMiddleware

Middleware para rastreamento de tempo de execução das requisições.

**Funcionalidades:**
- Define header `X-Request-Start` com timestamp de início
- Define header `X-Response-Time` com duração total da requisição
- Permite monitoramento de performance das operações

**Headers adicionados:**
- `X-Request-Start`: Timestamp do início da requisição em formato RFC3339Nano
- `X-Response-Time`: Duração total da requisição (ex: "15.234ms")

## ProcessInputMiddleware

Middleware que combina sanitização e validação de dados de entrada.

**Sanitização:**
- Remove espaços em branco extras do início e fim
- Converte emails para lowercase
- Remove caracteres especiais de CPF/CNPJ e telefone para validação
- Aplica formatação padrão após validação

**Validação:**
- **Nome**: Obrigatório, entre 2 e 255 caracteres
- **Email**: Obrigatório, formato válido
- **CPF/CNPJ**: Obrigatório, validação matemática dos dígitos verificadores
- **Telefone**: Obrigatório, formato brasileiro válido

**Aplicação:**
Aplicado apenas em rotas POST, PUT e PATCH que modificam dados.

## Estrutura de Resposta de Erro

Quando um middleware encontra um erro de validação, retorna:

```json
{
  "error": "Descrição do erro"
}
```

**Códigos de status HTTP:**
- `400 Bad Request`: Dados inválidos ou formato incorreto
- `500 Internal Server Error`: Erro interno de processamento

## Ordem de Execução

Na aplicação, os middlewares são executados na seguinte ordem:

1. **LoggingMiddleware** - Aplicado em todas as rotas da API
2. **TracingMiddleware** - Aplicado em todas as rotas da API  
3. **ProcessInputMiddleware** - Aplicado apenas na rota POST de criação de contatos

## Testes

Todos os middlewares possuem testes unitários que verificam:

- Funcionamento correto com dados válidos
- Tratamento adequado de dados inválidos
- Sanitização correta dos campos de entrada
- Validação de regras de negócio
- Geração correta de logs e headers de tracing

Para executar os testes:
```bash
go test -v ./internal/middleware/
```
