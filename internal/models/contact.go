package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mvcoladello/api-go-crm-contatos/internal/validators"
	"gorm.io/gorm"
)

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

func (c *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	if err := c.SanitizeAndValidate(); err != nil {
		return err
	}

	return
}

func (c *Contact) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := c.SanitizeAndValidate(); err != nil {
		return err
	}

	return
}

func (c *Contact) SanitizeAndValidate() error {
	c.Nome = validators.SanitizeName(c.Nome)
	c.Email = validators.SanitizeEmail(c.Email)
	c.CPFCNPJ = validators.SanitizeInput(c.CPFCNPJ)
	c.Telefone = validators.SanitizeInput(c.Telefone)

	if c.Nome == "" {
		return errors.New("nome é obrigatório")
	}

	if !validators.ValidateEmail(c.Email) {
		return errors.New("email inválido")
	}

	if !validators.ValidateDocument(c.CPFCNPJ) {
		return errors.New("CPF/CNPJ inválido")
	}

	if !validators.ValidateBrazilianPhone(c.Telefone) {
		return errors.New("telefone inválido")
	}

	c.CPFCNPJ = validators.FormatDocument(c.CPFCNPJ)
	c.Telefone = validators.FormatBrazilianPhone(c.Telefone)

	return nil
}

func (c *Contact) GetDocumentType() string {
	return validators.GetDocumentType(c.CPFCNPJ)
}

func (c *Contact) GetPhoneType() string {
	return validators.GetPhoneType(c.Telefone)
}

func (Contact) TableName() string {
	return "contacts"
}
