.PHONY: help build run dev stop clean seed test docker-build docker-up docker-down docker-logs

help: ## Mostra esta ajuda
	@echo "🚀 CRM Contatos - Comandos disponíveis:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Compila a aplicação
	@echo "🔨 Compilando aplicação..."
	go build -o main .

run: ## Executa a aplicação localmente
	@echo "🚀 Executando aplicação..."
	go run main.go

dev: ## Executa em modo desenvolvimento com hot-reload
	@echo "🔥 Executando em modo desenvolvimento..."
	air

stop: ## Para a aplicação
	@echo "🛑 Parando aplicação..."
	@pkill -f "go run main.go" || true
	@pkill -f "./main" || true

clean: ## Limpa arquivos de build
	@echo "🧹 Limpando arquivos..."
	rm -f main seed
	rm -f *.db

seed: ## Executa o seed de dados
	@echo "🌱 Executando seed de dados..."
	@./scripts/run-seed.sh

test: ## Executa os testes
	@echo "🧪 Executando testes..."
	go test -v ./...

test-validators: ## Executa apenas os testes dos validadores
	@echo "🧪 Executando testes dos validadores..."
	go test -v ./internal/validators/

# Quality Commands
lint: ## Executa linter
	@echo "🔍 Executando linter..."
	golangci-lint run

format: ## Formata o código
	@echo "✨ Formatando código..."
	gofmt -w .
	goimports -w .

# Performance & Security Tests
benchmark: ## Executa testes de performance
	@echo "⚡ Executando benchmarks..."
	go test -bench=. -benchmem ./...

race: ## Testa condições de corrida
	@echo "🏃 Testando race conditions..."
	go test -race ./...

coverage: ## Gera relatório de cobertura
	@echo "📊 Gerando relatório de cobertura..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Relatório gerado em: coverage.html"

test-integration: ## Executa testes de integração
	@echo "🔗 Executando testes de integração..."
	go test -tags=integration ./...

test-all: ## Executa todos os tipos de teste
	@echo "🧪 Executando todos os testes..."
	$(MAKE) test
	$(MAKE) benchmark
	$(MAKE) race
	$(MAKE) coverage

# Security
security-scan: ## Executa scan de segurança
	@echo "🔒 Executando scan de segurança..."
	gosec ./...

# Docker Commands
docker-build: ## Constrói a imagem Docker
	@echo "🐳 Construindo imagem Docker..."
	docker-compose build

docker-up: ## Sobe os containers em produção
	@echo "🚀 Subindo containers (produção)..."
	docker-compose up -d

docker-dev: ## Sobe os containers em desenvolvimento
	@echo "🔥 Subindo containers (desenvolvimento)..."
	docker-compose -f docker-compose.dev.yml up -d

docker-down: ## Para os containers
	@echo "🛑 Parando containers..."
	docker-compose down
	docker-compose -f docker-compose.dev.yml down

docker-logs: ## Mostra os logs dos containers
	@echo "📄 Logs dos containers..."
	docker-compose logs -f

docker-clean: ## Remove containers e volumes
	@echo "🧹 Limpando Docker..."
	docker-compose down -v
	docker-compose -f docker-compose.dev.yml down -v
	docker system prune -f

# Database Commands
db-migrate: ## Executa migração do banco
	@echo "📊 Executando migração..."
	go run -tags migrate scripts/migrate.go

db-migrate-down: ## Reverte última migração
	@echo "⬇️ Revertendo migração..."
	go run -tags migrate scripts/migrate.go -action=down

db-migrate-status: ## Mostra status das migrações
	@echo "📋 Status das migrações..."
	go run -tags migrate scripts/migrate.go -action=status

db-migrate-reset: ## Reseta todas as migrações
	@echo "🔄 Resetando migrações..."
	go run -tags migrate scripts/migrate.go -action=reset

db-reset: ## Reseta o banco de dados
	@echo "🗑️ Resetando banco de dados..."
	rm -f crm_contatos.db
	$(MAKE) db-migrate
	$(MAKE) seed

# API Documentation
swagger-gen: ## Gera documentação Swagger
	@echo "📚 Gerando documentação Swagger..."
	swag init -g main.go -o docs/

swagger-serve: ## Serve documentação Swagger
	@echo "🌐 Servindo documentação Swagger..."
	@echo "Acesse: http://localhost:3000/docs/"

# Production Commands
deploy: ## Deploy para produção
	@echo "🚀 Fazendo deploy..."
	$(MAKE) test
	$(MAKE) docker-build
	$(MAKE) docker-up

health-check: ## Verifica saúde da aplicação
	@echo "💚 Verificando saúde da aplicação..."
	curl -f http://localhost:3000/health || echo "❌ Aplicação não está respondendo"

# Install Commands
install-deps: ## Instala dependências de desenvolvimento
	@echo "📦 Instalando dependências..."
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/air-verse/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

install-swagger: ## Instala dependências do Swagger
	@echo "📚 Instalando Swagger..."
	go get -u github.com/gofiber/swagger
	go get -u github.com/swaggo/files
	go get -u github.com/swaggo/swag
