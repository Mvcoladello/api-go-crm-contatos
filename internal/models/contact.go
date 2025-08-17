package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Contact representa um contato no sistema CRM
type Contact struct {
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Nome      string         `json:"nome" gorm:"not null;size:255" validate:"required,min=2,max=255"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:255" validate:"required,email"`
	CPFCNPJ   string         `json:"cpf_cnpj" gorm:"uniqueIndex;not null;size:18" validate:"required"`
	Telefone  string         `json:"telefone" gorm:"not null;size:20" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// BeforeCreate é um hook do GORM que é executado antes de criar um registro
func (c *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}

// TableName define o nome da tabela no banco de dados
func (Contact) TableName() string {
	return "contacts"
}
