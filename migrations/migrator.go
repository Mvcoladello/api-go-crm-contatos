package migrations

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// MigrationRecord representa um registro de migração no banco
type MigrationRecord struct {
	ID          string    `gorm:"primaryKey"`
	Description string    `gorm:"not null"`
	AppliedAt   time.Time `gorm:"not null"`
}

// TableName define o nome da tabela para MigrationRecord
func (MigrationRecord) TableName() string {
	return "schema_migrations"
}

// Migrator gerencia as migrações
type Migrator struct {
	db *gorm.DB
}

// NewMigrator cria um novo migrator
func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

// createMigrationsTable cria a tabela de controle de migrações
func (m *Migrator) createMigrationsTable() error {
	return m.db.AutoMigrate(&MigrationRecord{})
}

// isMigrationApplied verifica se uma migração já foi aplicada
func (m *Migrator) isMigrationApplied(migrationID string) (bool, error) {
	var count int64
	err := m.db.Model(&MigrationRecord{}).Where("id = ?", migrationID).Count(&count).Error
	return count > 0, err
}

// recordMigration registra uma migração como aplicada
func (m *Migrator) recordMigration(migration Migration) error {
	record := MigrationRecord{
		ID:          migration.ID,
		Description: migration.Description,
		AppliedAt:   time.Now(),
	}
	return m.db.Create(&record).Error
}

// removeMigrationRecord remove o registro de uma migração
func (m *Migrator) removeMigrationRecord(migrationID string) error {
	return m.db.Where("id = ?", migrationID).Delete(&MigrationRecord{}).Error
}

// Up executa todas as migrações pendentes
func (m *Migrator) Up() error {
	// Criar tabela de controle de migrações
	if err := m.createMigrationsTable(); err != nil {
		return fmt.Errorf("erro ao criar tabela de migrações: %w", err)
	}

	migrations := GetMigrations()
	appliedCount := 0

	for _, migration := range migrations {
		applied, err := m.isMigrationApplied(migration.ID)
		if err != nil {
			return fmt.Errorf("erro ao verificar migração %s: %w", migration.ID, err)
		}

		if applied {
			log.Printf("Migração %s já aplicada, pulando...", migration.ID)
			continue
		}

		log.Printf("Aplicando migração %s: %s", migration.ID, migration.Description)

		// Executar migração em transação
		err = m.db.Transaction(func(tx *gorm.DB) error {
			if err := migration.Up(tx); err != nil {
				return fmt.Errorf("erro ao executar migração: %w", err)
			}

			// Registrar migração como aplicada
			record := MigrationRecord{
				ID:          migration.ID,
				Description: migration.Description,
				AppliedAt:   time.Now(),
			}
			return tx.Create(&record).Error
		})

		if err != nil {
			return fmt.Errorf("erro ao aplicar migração %s: %w", migration.ID, err)
		}

		appliedCount++
		log.Printf("Migração %s aplicada com sucesso", migration.ID)
	}

	if appliedCount == 0 {
		log.Println("Nenhuma migração pendente encontrada")
	} else {
		log.Printf("Total de %d migração(ões) aplicada(s) com sucesso", appliedCount)
	}

	return nil
}

// Down reverte a última migração aplicada
func (m *Migrator) Down() error {
	migrations := GetMigrations()

	// Buscar a última migração aplicada
	var lastRecord MigrationRecord
	err := m.db.Order("applied_at DESC").First(&lastRecord).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Nenhuma migração encontrada para reverter")
			return nil
		}
		return fmt.Errorf("erro ao buscar última migração: %w", err)
	}

	// Encontrar a migração correspondente
	var targetMigration *Migration
	for _, migration := range migrations {
		if migration.ID == lastRecord.ID {
			targetMigration = &migration
			break
		}
	}

	if targetMigration == nil {
		return fmt.Errorf("migração %s não encontrada na lista de migrações", lastRecord.ID)
	}

	log.Printf("Revertendo migração %s: %s", targetMigration.ID, targetMigration.Description)

	// Executar rollback em transação
	err = m.db.Transaction(func(tx *gorm.DB) error {
		if err := targetMigration.Down(tx); err != nil {
			return fmt.Errorf("erro ao executar rollback: %w", err)
		}

		// Remover registro da migração
		return tx.Where("id = ?", targetMigration.ID).Delete(&MigrationRecord{}).Error
	})

	if err != nil {
		return fmt.Errorf("erro ao reverter migração %s: %w", targetMigration.ID, err)
	}

	log.Printf("Migração %s revertida com sucesso", targetMigration.ID)
	return nil
}

// Status mostra o status das migrações
func (m *Migrator) Status() error {
	if err := m.createMigrationsTable(); err != nil {
		return fmt.Errorf("erro ao criar tabela de migrações: %w", err)
	}

	migrations := GetMigrations()

	fmt.Println("Status das Migrações:")
	fmt.Println("=====================")

	for _, migration := range migrations {
		applied, err := m.isMigrationApplied(migration.ID)
		if err != nil {
			return fmt.Errorf("erro ao verificar migração %s: %w", migration.ID, err)
		}

		status := "PENDENTE"
		if applied {
			status = "APLICADA"
		}

		fmt.Printf("%-30s %-10s %s\n", migration.ID, status, migration.Description)
	}

	return nil
}

// Reset remove todas as migrações e reaplica
func (m *Migrator) Reset() error {
	log.Println("Resetando banco de dados...")

	// Reverter todas as migrações
	for {
		var count int64
		if err := m.db.Model(&MigrationRecord{}).Count(&count).Error; err != nil {
			return fmt.Errorf("erro ao contar migrações: %w", err)
		}

		if count == 0 {
			break
		}

		if err := m.Down(); err != nil {
			return fmt.Errorf("erro ao reverter migração: %w", err)
		}
	}

	// Reaplicar todas as migrações
	return m.Up()
}
