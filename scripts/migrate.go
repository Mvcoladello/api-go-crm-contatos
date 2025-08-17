//go:build migrate
// +build migrate

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mvcoladello/api-go-crm-contatos/migrations"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		action = flag.String("action", "up", "Ação a ser executada: up, down, status, reset")
		dbPath = flag.String("db", "crm.db", "Caminho para o banco de dados SQLite")
		help   = flag.Bool("help", false, "Mostra esta ajuda")
	)
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	// Conectar ao banco de dados
	db, err := gorm.Open(sqlite.Open(*dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// Criar migrator
	migrator := migrations.NewMigrator(db)

	// Executar ação
	switch *action {
	case "up":
		if err := migrator.Up(); err != nil {
			log.Fatalf("Erro ao executar migrações: %v", err)
		}
	case "down":
		if err := migrator.Down(); err != nil {
			log.Fatalf("Erro ao reverter migração: %v", err)
		}
	case "status":
		if err := migrator.Status(); err != nil {
			log.Fatalf("Erro ao verificar status: %v", err)
		}
	case "reset":
		if err := migrator.Reset(); err != nil {
			log.Fatalf("Erro ao resetar banco: %v", err)
		}
	default:
		fmt.Printf("Ação inválida: %s\n", *action)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Ferramenta de Migração - CRM Contatos")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println("Uso: go run scripts/migrate.go [opções]")
	fmt.Println()
	fmt.Println("Opções:")
	fmt.Println("  -action string")
	fmt.Println("        Ação a ser executada: up, down, status, reset (padrão: up)")
	fmt.Println("  -db string")
	fmt.Println("        Caminho para o banco de dados SQLite (padrão: crm.db)")
	fmt.Println("  -help")
	fmt.Println("        Mostra esta ajuda")
	fmt.Println()
	fmt.Println("Ações:")
	fmt.Println("  up     - Aplica todas as migrações pendentes")
	fmt.Println("  down   - Reverte a última migração aplicada")
	fmt.Println("  status - Mostra o status de todas as migrações")
	fmt.Println("  reset  - Remove todas as migrações e reaplica")
	fmt.Println()
	fmt.Println("Exemplos:")
	fmt.Println("  go run scripts/migrate.go                    # Aplica migrações pendentes")
	fmt.Println("  go run scripts/migrate.go -action=status     # Mostra status")
	fmt.Println("  go run scripts/migrate.go -action=down       # Reverte última migração")
	fmt.Println("  go run scripts/migrate.go -action=reset      # Reset completo")
}
