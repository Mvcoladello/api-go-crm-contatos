# 📇 API CRM Contatos

![Go](https://img.shields.io/badge/Go-1.24.5-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2-00D9FF?style=for-the-badge&logo=fiber)
![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=for-the-badge&logo=sqlite)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)
![Swagger](https://img.shields.io/badge/Swagger-OpenAPI-85EA2D?style=for-the-badge&logo=swagger)

API REST moderna e eficiente para gerenciamento de contatos com validação robusta de documentos brasileiros (CPF/CNPJ), telefones, rate limiting e documentação completa.

## 🚀 Características Principais

- ✅ **CRUD Completo** - Criar, listar, buscar e deletar contatos
- ✅ **Validação Brasileira** - CPF, CNPJ e telefones brasileiros
- ✅ **Rate Limiting** - Proteção contra spam e ataques DDoS
- ✅ **Swagger/OpenAPI** - Documentação interativa da API
- ✅ **Migrations** - Sistema explícito de migração de banco
- ✅ **Segurança** - Sanitização de entrada e prevenção XSS
- ✅ **Performance** - Framework Fiber de alta performance
- ✅ **Banco de Dados** - SQLite com GORM ORM
- ✅ **Docker Ready** - Containerização completa com Nginx
- ✅ **Testes Abrangentes** - Testes unitários, integração e performance
- ✅ **Hot Reload** - Desenvolvimento com recarga automática
- ✅ **Middleware** - Logging, CORS, rate limiting e validação
- ✅ **Formatação Automática** - Dados brasileiros formatados

## 📋 Funcionalidades

### 🔍 Operações de Contato
- Listar todos os contatos com paginação
- Buscar contato específico por ID (UUID)
- Criar novo contato com validação completa
- Deletar contato existente
- Prevenção de duplicatas (email e CPF/CNPJ únicos)

### 🛡️ Validações e Segurança
- **CPF**: Validação matemática + formatação automática
- **CNPJ**: Validação matemática + formatação automática  
- **Telefone**: Validação de DDD + formatação brasileira
- **Email**: Validação de formato + unicidade
- **Rate Limiting**: Proteção contra abuso (configurável por endpoint)
- **Sanitização**: Prevenção XSS e limpeza de dados
- **Middleware**: Logging de requisições e tratamento de erros

### 📚 Documentação
- **Swagger UI**: Interface interativa para testar endpoints
- **OpenAPI 3.0**: Especificação completa da API
- **Exemplos**: Casos de uso documentados
- **Schemas**: Modelos de dados validados

### 🗄️ Banco de Dados
- **Migrations**: Sistema de versionamento de schema
- **Rollback**: Capacidade de reverter alterações
- **Índices**: Otimização de consultas
- **Constraints**: Integridade referencial

## 🏗️ Arquitetura

```
📦 api-go-crm-contatos/
├── 📚 docs/                       # Documentação Swagger/OpenAPI
│   └── swagger.yaml               # Especificação OpenAPI
├── 🗄️ migrations/                 # Sistema de migração
│   ├── migrations.go              # Definições das migrações
│   ├── migrator.go               # Engine de migração
│   └── migrations_test.go        # Testes das migrações
├── 📦 scripts/                    # Scripts utilitários
│   ├── migrate.go                # CLI para migrações
│   ├── seed.go                   # Dados de exemplo
│   └── run-seed.sh              # Script de seed
├── 🏠 internal/                   # Código da aplicação
│   ├── 🎮 handlers/              # Controllers da API
│   │   ├── contact_handler.go    # Endpoints de contatos
│   │   ├── contact_handler_test.go # Testes dos handlers
│   │   └── routes.go             # Definição de rotas
│   ├── ⚙️ middleware/            # Middlewares personalizados
│   │   ├── rate_limit.go         # Rate limiting
│   │   ├── rate_limit_test.go    # Testes do rate limiting
│   │   ├── logging.go            # Middleware de logging
│   │   ├── process_input.go      # Processamento de entrada
│   │   ├── sanitize.go           # Sanitização de dados
│   │   ├── validation.go         # Validações
│   │   └── middleware_test.go    # Testes dos middlewares
│   ├── 📊 models/                # Modelos de dados
│   │   ├── contact.go            # Modelo de contato
│   │   └── responses.go          # Modelos de resposta
│   ├── 🔧 services/              # Lógica de negócio
│   │   └── contact_service.go    # Serviços de contato
│   ├── 🛠️ utils/                 # Utilitários
│   │   └── responses.go          # Helpers de resposta
│   └── ✅ validators/            # Validadores específicos
│       ├── cpf.go                # Validação de CPF
│       ├── cnpj.go               # Validação de CNPJ
│       ├── phone.go              # Validação de telefone
│       ├── sanitizer.go          # Sanitização
│       ├── examples.go           # Exemplos de uso
│       └── validators_test.go    # Testes dos validadores
├── 🐳 docker-compose.yml          # Orquestração de containers
├── 🐳 docker-compose.dev.yml      # Ambiente de desenvolvimento
├── 🐳 Dockerfile                  # Imagem da aplicação
├── 🔧 Makefile                    # Comandos de automação
├── 📄 nginx.conf                  # Configuração do proxy reverso
├── 🚀 main.go                     # Ponto de entrada da aplicação
├── 📋 go.mod                      # Dependências do projeto
├── 🧪 integration_test.go         # Testes de integração
└── 📖 README.md                   # Esta documentação
```

## 🚀 Quick Start

### Opção 1: Docker (Recomendado)
```bash
# Clonar repositório
git clone https://github.com/mvcoladello/api-go-crm-contatos.git
cd api-go-crm-contatos

# Executar com Docker
make docker-up

# A API estará disponível em http://localhost:8080
# Documentação Swagger em http://localhost:8080/docs/
```

### Opção 2: Execução Local
```bash
# Instalar dependências
make install-deps

# Executar migrações
make db-migrate

# Popular banco com dados de exemplo
make seed

# Executar aplicação
make dev

# A API estará disponível em http://localhost:3000
# Documentação Swagger em http://localhost:3000/docs/
```

## 📚 Documentação da API

### Swagger UI
Acesse a documentação interativa em:
- **Local**: http://localhost:3000/docs/
- **Docker**: http://localhost:8080/docs/

### Endpoints Principais

| Método | Endpoint | Descrição | Rate Limit |
|--------|----------|-----------|------------|
| GET | `/health` | Health check | Sem limite |
| GET | `/contacts` | Listar contatos | 100 req/s |
| POST | `/contacts` | Criar contato | 100 req/s |
| GET | `/contacts/{id}` | Buscar contato | 100 req/s |
| PUT | `/contacts/{id}` | Atualizar contato | 100 req/s |
| DELETE | `/contacts/{id}` | Deletar contato | 100 req/s |

## 🚦 Endpoints da API

### Base URL
```
http://localhost:3000/api/v1
```

### 📋 Listar Contatos
```http
GET /contatos
```

**Resposta (200):**
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

### 🔍 Buscar Contato por ID
```http
GET /contatos/:id
```

**Parâmetros:**
- `id` (UUID): Identificador único do contato

**Resposta (200):**
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

### ✏️ Criar Contato
```http
POST /contatos
```

**Corpo da Requisição:**
```json
{
  "nome": "João Silva",
  "email": "joao@example.com",
  "cpf_cnpj": "12345678900",
  "telefone": "11999999999"
}
```

**Resposta (201):**
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

### 🗑️ Deletar Contato
```http
DELETE /contatos/:id
```

**Resposta (200):**
```json
{
  "message": "Contato deletado com sucesso"
}
```

### 💓 Health Check
```http
GET /health
```

**Resposta (200):**
```json
{
  "status": "healthy",
  "timestamp": "2025-08-17T10:00:00Z"
}
```

## 🔧 Instalação e Execução

### 📋 Pré-requisitos
- **Go 1.24.5+**
- **Docker** (opcional)
- **Make** (opcional, para comandos automatizados)

### 🚀 Execução Local

#### 1. Clone o repositório
```bash
git clone https://github.com/mvcoladello/api-go-crm-contatos.git
cd api-go-crm-contatos
```

#### 2. Instale as dependências
```bash
go mod download
```

#### 3. Execute a aplicação
```bash
# Usando Go
go run main.go

# Ou usando Make
make run

# Para desenvolvimento com hot-reload
make dev
```

#### 4. Acesse a API
```
http://localhost:3000
```

### 🐳 Execução com Docker

#### Produção
```bash
# Inicie todos os serviços
docker-compose up -d

# Ou usando Make
make docker-up
```

#### Desenvolvimento
```bash
# Ambiente de desenvolvimento
docker-compose -f docker-compose.dev.yml up -d

# Ou usando Make
make docker-dev
```

## 🗄️ Sistema de Migrations

### Comandos de Migration
```bash
# Executar todas as migrações pendentes
make db-migrate

# Verificar status das migrações
make db-migrate-status

# Reverter última migração
make db-migrate-down

# Resetar todas as migrações (cuidado!)
make db-migrate-reset
```

### Estrutura das Migrations
As migrations estão organizadas no diretório `migrations/` e incluem:
- `001_create_contacts_table` - Criação da tabela de contatos
- `002_add_indexes_to_contacts` - Adição de índices para performance

### Adicionando Nova Migration
```go
// Em migrations/migrations.go
{
    ID:          "003_new_migration",
    Description: "Descrição da nova migração",
    Up:          newMigrationUp,
    Down:        newMigrationDown,
}
```

## 🛡️ Rate Limiting

### Configuração por Endpoint
- **Endpoints gerais**: 100 req/s com burst de 200
- **Health check**: Sem limite
- **Endpoints sensíveis**: 5 req/s com burst de 10

### Configuração Customizada
```go
// Rate limit personalizado
app.Use(middleware.RateLimiter(middleware.RateLimitConfig{
    Rate:  rate.Limit(50),
    Burst: 100,
    KeyGenerator: func(c *fiber.Ctx) string {
        return c.IP() + ":" + c.Get("User-Agent")
    },
}))
```

### Monitoramento
```bash
# Verificar estatísticas de rate limiting
curl http://localhost:3000/admin/rate-limit/stats
```

## 🧪 Testes

### Executar Todos os Testes
```bash
# Testes unitários
make test

# Testes com cobertura
make coverage

# Testes de performance (benchmarks)
make benchmark

# Testes de condições de corrida
make race

# Todos os tipos de teste
make test-all
```

### Estrutura de Testes
- **Unitários**: Cada módulo possui seu arquivo `*_test.go`
- **Integração**: `integration_test.go` testa fluxos completos
- **Migrations**: `migrations/migrations_test.go` testa sistema de migração
- **Rate Limiting**: `internal/middleware/rate_limit_test.go`

### Cobertura de Testes
```bash
# Gerar relatório HTML de cobertura
make coverage
open coverage.html
```

## 🔧 Comandos do Makefile

### Desenvolvimento
```bash
make help          # Lista todos os comandos disponíveis
make dev           # Executa em modo desenvolvimento (hot-reload)
make build         # Compila a aplicação
make run           # Executa a aplicação compilada
```

### Banco de Dados
```bash
make db-migrate    # Executa migrações
make db-reset      # Reseta banco e reaplica migrações
make seed          # Popula banco com dados de exemplo
```

### Docker
```bash
make docker-build  # Constrói imagem Docker
make docker-up     # Sobe containers
make docker-down   # Para containers
make docker-logs   # Visualiza logs
```

### Qualidade de Código
```bash
make lint          # Executa linter
make format        # Formata código
make security-scan # Scan de segurança
```

### Documentação
```bash
make swagger-gen   # Gera documentação Swagger
make swagger-serve # Serve documentação localmente
```

## 🔍 Validadores Brasileiros

### 📄 CPF (Cadastro de Pessoa Física)
- Validação matemática dos dígitos verificadores
- Formatação automática: `XXX.XXX.XXX-XX`
- Rejeita CPFs inválidos conhecidos (ex: 111.111.111-11)
- Aceita entrada com ou sem formatação

### 🏢 CNPJ (Cadastro Nacional da Pessoa Jurídica)
- Validação matemática dos dígitos verificadores
- Formatação automática: `XX.XXX.XXX/XXXX-XX`
- Rejeita CNPJs inválidos conhecidos
- Aceita entrada com ou sem formatação

### 📞 Telefone Brasileiro
- Suporte a telefones fixos (10 dígitos) e celulares (11 dígitos)
- Validação de DDDs brasileiros válidos
- Formatação automática: `(XX) XXXXX-XXXX` ou `(XX) XXXX-XXXX`
- Suporte opcional ao código do país (+55)

### 📧 Email
- Validação de formato padrão
- Verificação de unicidade no banco
- Sanitização automática

## 🛡️ Segurança

### Sanitização de Dados
- **XSS Prevention**: Escape de caracteres HTML perigosos
- **Input Cleaning**: Remoção de caracteres de controle
- **Data Validation**: Validação rigorosa de todos os campos
- **SQL Injection**: Proteção via GORM ORM

### Middleware de Segurança
- **Logging**: Registro de todas as requisições
- **CORS**: Configuração de Cross-Origin Resource Sharing
- **Error Handling**: Tratamento centralizado de erros
- **Input Processing**: Processamento seguro de entrada

## 🗄️ Banco de Dados

### SQLite
- Banco embarcado de alta performance
- Arquivo: `crm_contatos.db` (criado automaticamente)
- Migrações automáticas via GORM
- Backup simples (arquivo único)

### Estrutura da Tabela `contacts`
```sql
CREATE TABLE contacts (
    id UUID PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    cpf_cnpj VARCHAR(18) UNIQUE NOT NULL,
    telefone VARCHAR(15) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);
```

## 🔄 Middlewares

### Logging
- Registro detalhado de requisições HTTP
- Informações de IP, método, rota e tempo de resposta
- Logs estruturados para análise

### Validação
- Validação automática de dados de entrada
- Sanitização preventiva contra XSS
- Formatação automática de dados brasileiros

### CORS
- Configuração flexível para desenvolvimento e produção
- Suporte a diferentes origens e métodos

### Rate Limiting
- Limitação de requisições por IP e por endpoint
- Proteção contra ataques de força bruta e DDoS
- Configuração de janelas de tempo e limites personalizados

## 🐳 Docker e Deploy

### Estrutura de Containers
- **API Container**: Aplicação Go otimizada
- **Nginx Container**: Proxy reverso e balanceamento
- **Volume Persistence**: Dados persistidos entre deploys

### Características do Deploy
- **Health Checks**: Monitoramento automático de saúde
- **Auto Restart**: Reinicialização automática em falhas
- **SSL Ready**: Configuração preparada para HTTPS
- **Multi-Stage Build**: Imagens otimizadas

### Nginx
- Proxy reverso para a API
- Configuração de SSL/TLS
- Balanceamento de carga (se necessário)
- Cache de recursos estáticos

## 📈 Performance

### Otimizações
- **Fiber Framework**: Framework web de alta performance
- **SQLite**: Banco embarcado otimizado
- **Connection Pooling**: Pool de conexões eficiente
- **Minimal Dependencies**: Dependências mínimas e focadas

### Métricas
- Tempo de resposta médio: < 10ms
- Throughput: 1000+ req/s (dependendo do hardware)
- Uso de memória: ~20MB em produção

## 🤝 Contribuição

1. **Fork** o projeto
2. **Clone** sua fork
3. **Crie** uma branch para sua feature
4. **Commit** suas mudanças
5. **Push** para a branch
6. **Abra** um Pull Request

### Padrões de Código
- Siga as convenções Go padrão
- Execute `go fmt` antes de commitar
- Mantenha cobertura de testes > 80%
- Documente funções públicas

## 🐛 Resolução de Problemas

### Problemas Comuns

#### Erro de Conexão com Banco
```bash
# Verifique permissões do diretório
chmod 755 .
# Recrie o banco
rm -f crm_contatos.db && go run main.go
```

#### Porta em Uso
```bash
# Verifique processos na porta 3000
lsof -i :3000
# Mate o processo se necessário
kill -9 <PID>
```

#### Dependências Desatualizadas
```bash
# Atualize dependências
go mod tidy
go mod download
```

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 📞 Contato

**Desenvolvedor**: Marcos Coladello  
**Email**: marcos@coladello.com.br  
**GitHub**: [@mvcoladello](https://github.com/mvcoladello)

---

<div align="center">

**⭐ Se este projeto foi útil, considere dar uma estrela!**

Made with ❤️ in Brazil 🇧🇷

</div>
