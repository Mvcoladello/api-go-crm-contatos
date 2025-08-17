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

test-coverage: ## Executa testes com cobertura
	@echo "📊 Executando testes com cobertura..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-coverage-validators: ## Executa testes dos validadores com cobertura
	@echo "📊 Executando testes dos validadores com cobertura..."
	go test -v -coverprofile=validators_coverage.out ./internal/validators/
	go tool cover -html=validators_coverage.out -o validators_coverage.html

bench-validators: ## Executa benchmarks dos validadores
	@echo "⚡ Executando benchmarks dos validadores..."
	go test -bench=. ./internal/validators/

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

db-reset: ## Reseta o banco de dados
	@echo "🗑️ Resetando banco de dados..."
	rm -f crm_contatos.db
	$(MAKE) db-migrate
	$(MAKE) seed

# Quality Commands
lint: ## Executa linter
	@echo "🔍 Executando linter..."
	golangci-lint run

format: ## Formata o código
	@echo "✨ Formatando código..."
	gofmt -w .
	goimports -w .

# Install Commands
install-deps: ## Instala dependências de desenvolvimento
	@echo "📦 Instalando dependências..."
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/air-verse/air@latest

# Production Commands
deploy: ## Deploy para produção
	@echo "🚀 Fazendo deploy..."
	$(MAKE) test
	$(MAKE) docker-build
	$(MAKE) docker-up

health-check: ## Verifica saúde da aplicação
	@echo "💚 Verificando saúde da aplicação..."
	curl -f http://localhost:3000/health || echo "❌ Aplicação não está respondendo"
