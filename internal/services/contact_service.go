package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
	"gorm.io/gorm"
)

type ContactService struct {
	db *gorm.DB
}

func NewContactService(db *gorm.DB) *ContactService {
	return &ContactService{
		db: db,
	}
}

func (s *ContactService) GetAllContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	result := s.db.Find(&contacts)
	return contacts, result.Error
}

func (s *ContactService) GetContactByID(id uuid.UUID) (*models.Contact, error) {
	var contact models.Contact
	result := s.db.First(&contact, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("contato não encontrado")
	}

	return &contact, result.Error
}

func (s *ContactService) CreateContact(contact *models.Contact) error {
	result := s.db.Create(contact)
	return result.Error
}

func (s *ContactService) UpdateContact(id uuid.UUID, updateData *models.Contact) (*models.Contact, error) {
	var contact models.Contact
	result := s.db.First(&contact, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("contato não encontrado")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	// Atualiza apenas os campos não vazios
	if updateData.Nome != "" {
		contact.Nome = updateData.Nome
	}
	if updateData.Email != "" {
		contact.Email = updateData.Email
	}
	if updateData.CPFCNPJ != "" {
		contact.CPFCNPJ = updateData.CPFCNPJ
	}
	if updateData.Telefone != "" {
		contact.Telefone = updateData.Telefone
	}

	result = s.db.Save(&contact)
	return &contact, result.Error
}

func (s *ContactService) DeleteContact(id uuid.UUID) error {
	result := s.db.Delete(&models.Contact{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return errors.New("contato não encontrado")
	}

	return result.Error
}
