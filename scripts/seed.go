package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("🌱 Iniciando seed de dados...")

	db, err := gorm.Open(sqlite.Open("crm_contatos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados:", err)
	}

	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatal("Falha na migração:", err)
	}

	var count int64
	db.Model(&models.Contact{}).Count(&count)
	if count > 0 {
		fmt.Printf("⚠️  Banco já possui %d contatos. Deseja limpar? (y/N): ", count)
		var response string
		fmt.Scanln(&response)
		if response == "y" || response == "Y" {
			db.Exec("DELETE FROM contacts")
			fmt.Println("🗑️  Dados anteriores removidos")
		} else {
			fmt.Println("ℹ️  Adicionando novos contatos aos existentes")
		}
	}

	contacts := []models.Contact{
		{
			ID:       uuid.New(),
			Nome:     "João Silva Santos",
			Email:    "joao.silva@email.com",
			CPFCNPJ:  "12345678901",
			Telefone: "11987654321",
		},
		{
			ID:       uuid.New(),
			Nome:     "Maria Oliveira Costa",
			Email:    "maria.oliveira@email.com",
			CPFCNPJ:  "98765432100",
			Telefone: "11876543210",
		},
		{
			ID:       uuid.New(),
			Nome:     "Empresa Tech Solutions LTDA",
			Email:    "contato@techsolutions.com.br",
			CPFCNPJ:  "12345678000195",
			Telefone: "1133334444",
		},
		{
			ID:       uuid.New(),
			Nome:     "Ana Paula Ferreira",
			Email:    "ana.ferreira@email.com",
			CPFCNPJ:  "11122233344",
			Telefone: "11912345678",
		},
		{
			ID:       uuid.New(),
			Nome:     "Carlos Eduardo Lima",
			Email:    "carlos.lima@email.com",
			CPFCNPJ:  "55566677788",
			Telefone: "11998887777",
		},
		{
			ID:       uuid.New(),
			Nome:     "Inovação Digital EIRELI",
			Email:    "comercial@inovacaodigital.com.br",
			CPFCNPJ:  "98765432000123",
			Telefone: "1144445555",
		},
		{
			ID:       uuid.New(),
			Nome:     "Fernanda Santos Rocha",
			Email:    "fernanda.rocha@email.com",
			CPFCNPJ:  "22233344455",
			Telefone: "11955554444",
		},
		{
			ID:       uuid.New(),
			Nome:     "Pedro Henrique Alves",
			Email:    "pedro.alves@email.com",
			CPFCNPJ:  "33344455566",
			Telefone: "11966665555",
		},
		{
			ID:       uuid.New(),
			Nome:     "Consultoria Empresarial S/A",
			Email:    "atendimento@consultoriaempresarial.com.br",
			CPFCNPJ:  "11223344000167",
			Telefone: "1155556666",
		},
		{
			ID:       uuid.New(),
			Nome:     "Beatriz Cardoso Mendes",
			Email:    "beatriz.mendes@email.com",
			CPFCNPJ:  "44455566677",
			Telefone: "11977778888",
		},
	}

	fmt.Printf("📝 Inserindo %d contatos...\n", len(contacts))

	for i, contact := range contacts {
		if err := db.Create(&contact).Error; err != nil {
			log.Printf("❌ Erro ao inserir contato %d (%s): %v", i+1, contact.Nome, err)
		} else {
			fmt.Printf("✅ %d. %s (%s)\n", i+1, contact.Nome, contact.Email)
		}
		time.Sleep(100 * time.Millisecond)
	}

	var finalCount int64
	db.Model(&models.Contact{}).Count(&finalCount)

	fmt.Printf("\n🎉 Seed concluído! Total de contatos no banco: %d\n", finalCount)
	fmt.Println("📊 Para visualizar os dados, execute: curl http://localhost:3000/api/v1/contatos")
}
