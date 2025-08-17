package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
	"github.com/mvcoladello/api-go-crm-contatos/internal/services"
)

type ContactHandler struct {
	contactService *services.ContactService
}

func NewContactHandler(contactService *services.ContactService) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
	}
}

// GetContacts lista todos os contatos
// GET /contatos
func (h *ContactHandler) GetContacts(c *fiber.Ctx) error {
	contacts, err := h.contactService.GetAllContacts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":  contacts,
		"total": len(contacts),
	})
}

// GetContact busca um contato por ID
// GET /contatos/:id
func (h *ContactHandler) GetContact(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "ID inválido",
			"details": "O ID deve ser um UUID válido",
		})
	}

	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		if err.Error() == "contato não encontrado" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Contato não encontrado",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": contact,
	})
}

// CreateContact cria um novo contato
// POST /contatos
func (h *ContactHandler) CreateContact(c *fiber.Ctx) error {
	var contact models.Contact

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Dados inválidos no corpo da requisição",
			"details": err.Error(),
		})
	}

	if err := h.contactService.CreateContact(&contact); err != nil {
		if err.Error() == "UNIQUE constraint failed: contacts.email" ||
			err.Error() == "UNIQUE constraint failed: contacts.cpf_cnpj" {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"error": "Email ou CPF/CNPJ já cadastrado",
			})
		}

		// Verifica se é erro de validação
		if err.Error() == "nome é obrigatório" ||
			err.Error() == "email inválido" ||
			err.Error() == "CPF/CNPJ inválido" ||
			err.Error() == "telefone inválido" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error":   "Dados inválidos",
				"details": err.Error(),
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Contato criado com sucesso",
		"data":    contact,
	})
}

// DeleteContact deleta um contato
// DELETE /contatos/:id
func (h *ContactHandler) DeleteContact(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "ID inválido",
			"details": "O ID deve ser um UUID válido",
		})
	}

	if err := h.contactService.DeleteContact(id); err != nil {
		if err.Error() == "contato não encontrado" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Contato não encontrado",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Contato deletado com sucesso",
	})
}
