package migrations

import (
	"gorm.io/gorm"
)

// Migration representa uma migração de banco de dados
type Migration struct {
	ID          string
	Description string
	Up          func(*gorm.DB) error
	Down        func(*gorm.DB) error
}

// GetMigrations retorna todas as migrações na ordem correta
func GetMigrations() []Migration {
	return []Migration{
		{
			ID:          "001_create_contacts_table",
			Description: "Criar tabela de contatos",
			Up:          createContactsTable,
			Down:        dropContactsTable,
		},
		{
			ID:          "002_add_indexes_to_contacts",
			Description: "Adicionar índices à tabela de contatos",
			Up:          addContactsIndexes,
			Down:        dropContactsIndexes,
		},
	}
}

// createContactsTable cria a tabela de contatos
func createContactsTable(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS contacts (
			id TEXT PRIMARY KEY,
			nome TEXT NOT NULL,
			email TEXT NOT NULL,
			telefone TEXT NOT NULL,
			cpf_cnpj TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			deleted_at DATETIME
		)
	`).Error
}

// dropContactsTable remove a tabela de contatos
func dropContactsTable(db *gorm.DB) error {
	return db.Exec("DROP TABLE IF EXISTS contacts").Error
}

// addContactsIndexes adiciona índices à tabela de contatos
func addContactsIndexes(db *gorm.DB) error {
	// Índice único para email
	if err := db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_contacts_email ON contacts(email) WHERE deleted_at IS NULL").Error; err != nil {
		return err
	}

	// Índice único para CPF/CNPJ
	if err := db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_contacts_cpf_cnpj ON contacts(cpf_cnpj) WHERE deleted_at IS NULL").Error; err != nil {
		return err
	}

	// Índice para nome (para buscas)
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_contacts_nome ON contacts(nome)").Error; err != nil {
		return err
	}

	// Índice para created_at (para ordenação)
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_contacts_created_at ON contacts(created_at)").Error; err != nil {
		return err
	}

	return nil
}

// dropContactsIndexes remove os índices da tabela de contatos
func dropContactsIndexes(db *gorm.DB) error {
	indexes := []string{
		"idx_contacts_email",
		"idx_contacts_cpf_cnpj",
		"idx_contacts_nome",
		"idx_contacts_created_at",
	}

	for _, index := range indexes {
		if err := db.Exec("DROP INDEX IF EXISTS " + index).Error; err != nil {
			return err
		}
	}

	return nil
}
