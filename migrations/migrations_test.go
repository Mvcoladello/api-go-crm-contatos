package migrations

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao conectar com banco de teste: %v", err)
	}
	return db
}

func TestMigrator_Up(t *testing.T) {
	db := setupTestDB(t)
	migrator := NewMigrator(db)

	err := migrator.Up()
	if err != nil {
		t.Fatalf("Erro ao executar migrações: %v", err)
	}

	// Verificar se a tabela de contatos foi criada
	if !db.Migrator().HasTable("contacts") {
		t.Error("Tabela 'contacts' não foi criada")
	}

	// Verificar se a tabela de migrações foi criada
	if !db.Migrator().HasTable("schema_migrations") {
		t.Error("Tabela 'schema_migrations' não foi criada")
	}

	// Verificar se as migrações foram registradas
	var count int64
	db.Model(&MigrationRecord{}).Count(&count)
	expectedMigrations := len(GetMigrations())

	if int(count) != expectedMigrations {
		t.Errorf("Esperado %d migrações registradas, obtido %d", expectedMigrations, count)
	}
}

func TestMigrator_UpIdempotent(t *testing.T) {
	db := setupTestDB(t)
	migrator := NewMigrator(db)

	// Executar migrações duas vezes
	err := migrator.Up()
	if err != nil {
		t.Fatalf("Erro na primeira execução: %v", err)
	}

	err = migrator.Up()
	if err != nil {
		t.Fatalf("Erro na segunda execução: %v", err)
	}

	// Verificar se não duplicou registros
	var count int64
	db.Model(&MigrationRecord{}).Count(&count)
	expectedMigrations := len(GetMigrations())

	if int(count) != expectedMigrations {
		t.Errorf("Migrações não são idempotentes. Esperado %d registros, obtido %d", expectedMigrations, count)
	}
}

func TestMigrator_Down(t *testing.T) {
	db := setupTestDB(t)
	migrator := NewMigrator(db)

	// Primeiro aplicar migrações
	err := migrator.Up()
	if err != nil {
		t.Fatalf("Erro ao aplicar migrações: %v", err)
	}

	initialCount := int64(0)
	db.Model(&MigrationRecord{}).Count(&initialCount)

	// Reverter uma migração
	err = migrator.Down()
	if err != nil {
		t.Fatalf("Erro ao reverter migração: %v", err)
	}

	// Verificar se o número de migrações diminuiu
	var finalCount int64
	db.Model(&MigrationRecord{}).Count(&finalCount)

	if finalCount != initialCount-1 {
		t.Errorf("Migração não foi revertida corretamente. Antes: %d, Depois: %d", initialCount, finalCount)
	}
}

func TestMigrator_Status(t *testing.T) {
	db := setupTestDB(t)
	migrator := NewMigrator(db)

	// Testar status sem migrações
	err := migrator.Status()
	if err != nil {
		t.Fatalf("Erro ao verificar status inicial: %v", err)
	}

	// Aplicar migrações e testar status novamente
	err = migrator.Up()
	if err != nil {
		t.Fatalf("Erro ao aplicar migrações: %v", err)
	}

	err = migrator.Status()
	if err != nil {
		t.Fatalf("Erro ao verificar status após migrações: %v", err)
	}
}

func TestMigrator_Reset(t *testing.T) {
	db := setupTestDB(t)
	migrator := NewMigrator(db)

	// Aplicar migrações
	err := migrator.Up()
	if err != nil {
		t.Fatalf("Erro ao aplicar migrações: %v", err)
	}

	// Verificar se existe dados
	var count int64
	db.Model(&MigrationRecord{}).Count(&count)
	if count == 0 {
		t.Fatal("Nenhuma migração foi aplicada antes do reset")
	}

	// Executar reset
	err = migrator.Reset()
	if err != nil {
		t.Fatalf("Erro ao executar reset: %v", err)
	}

	// Verificar se as migrações foram reaplicadas
	db.Model(&MigrationRecord{}).Count(&count)
	expectedMigrations := len(GetMigrations())

	if int(count) != expectedMigrations {
		t.Errorf("Reset não reaplicou todas as migrações. Esperado %d, obtido %d", expectedMigrations, count)
	}
}

func TestMigrations_Structure(t *testing.T) {
	migrations := GetMigrations()

	if len(migrations) == 0 {
		t.Fatal("Nenhuma migração encontrada")
	}

	// Verificar se todas as migrações têm ID e descrição
	for i, migration := range migrations {
		if migration.ID == "" {
			t.Errorf("Migração %d não tem ID", i)
		}

		if migration.Description == "" {
			t.Errorf("Migração %d não tem descrição", i)
		}

		if migration.Up == nil {
			t.Errorf("Migração %s não tem função Up", migration.ID)
		}

		if migration.Down == nil {
			t.Errorf("Migração %s não tem função Down", migration.ID)
		}
	}
}

func TestCreateContactsTable(t *testing.T) {
	db := setupTestDB(t)

	err := createContactsTable(db)
	if err != nil {
		t.Fatalf("Erro ao criar tabela de contatos: %v", err)
	}

	if !db.Migrator().HasTable("contacts") {
		t.Error("Tabela 'contacts' não foi criada")
	}

	// Verificar colunas
	expectedColumns := []string{"id", "nome", "email", "telefone", "cpf_cnpj", "created_at", "updated_at", "deleted_at"}
	for _, col := range expectedColumns {
		if !db.Migrator().HasColumn("contacts", col) {
			t.Errorf("Coluna '%s' não encontrada na tabela contacts", col)
		}
	}
}

func TestDropContactsTable(t *testing.T) {
	db := setupTestDB(t)

	// Criar tabela primeiro
	err := createContactsTable(db)
	if err != nil {
		t.Fatalf("Erro ao criar tabela: %v", err)
	}

	// Verificar se existe
	if !db.Migrator().HasTable("contacts") {
		t.Fatal("Tabela não foi criada")
	}

	// Dropar tabela
	err = dropContactsTable(db)
	if err != nil {
		t.Fatalf("Erro ao dropar tabela: %v", err)
	}

	// Verificar se foi removida
	if db.Migrator().HasTable("contacts") {
		t.Error("Tabela 'contacts' ainda existe após drop")
	}
}

func TestAddContactsIndexes(t *testing.T) {
	db := setupTestDB(t)

	// Criar tabela primeiro
	err := createContactsTable(db)
	if err != nil {
		t.Fatalf("Erro ao criar tabela: %v", err)
	}

	// Adicionar índices
	err = addContactsIndexes(db)
	if err != nil {
		t.Fatalf("Erro ao adicionar índices: %v", err)
	}

	// Verificar se os índices foram criados (SQLite específico)
	var count int64
	err = db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='index' AND tbl_name='contacts'").Scan(&count).Error
	if err != nil {
		t.Fatalf("Erro ao verificar índices: %v", err)
	}

	// Deve ter pelo menos alguns índices personalizados
	if count < 3 {
		t.Errorf("Esperado pelo menos 3 índices, encontrado %d", count)
	}
}
