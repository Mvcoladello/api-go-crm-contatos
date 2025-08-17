package handlers

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
	"github.com/mvcoladello/api-go-crm-contatos/internal/services"
	"github.com/mvcoladello/api-go-crm-contatos/internal/utils"
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
		return utils.SendInternalServerError(c, "Erro interno do servidor", err.Error())
	}

	return utils.SendSuccessResponseWithTotal(c, http.StatusOK, "", contacts, len(contacts))
}

// GetContact busca um contato por ID
// GET /contatos/:id
func (h *ContactHandler) GetContact(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return utils.SendBadRequestError(c, "ID deve ser um UUID válido")
	}

	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		if err.Error() == "contato não encontrado" {
			return utils.SendNotFoundError(c, "Contato não encontrado")
		}
		return utils.SendInternalServerError(c, "Erro interno do servidor", err.Error())
	}

	return utils.SendSuccessResponse(c, http.StatusOK, "", contact)
}

// CreateContact cria um novo contato
// POST /contatos
func (h *ContactHandler) CreateContact(c *fiber.Ctx) error {
	var contact models.Contact

	if err := c.BodyParser(&contact); err != nil {
		return utils.SendBadRequestError(c, "Dados inválidos no corpo da requisição")
	}

	if err := h.contactService.CreateContact(&contact); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return utils.SendConflictError(c, "Email ou CPF/CNPJ já cadastrado")
		}

		return utils.SendInternalServerError(c, "Erro interno do servidor", err.Error())
	}

	return utils.SendSuccessResponse(c, http.StatusCreated, "Contato criado com sucesso", contact)
}

// DeleteContact deleta um contato
// DELETE /contatos/:id
func (h *ContactHandler) DeleteContact(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return utils.SendBadRequestError(c, "ID deve ser um UUID válido")
	}

	if err := h.contactService.DeleteContact(id); err != nil {
		if err.Error() == "contato não encontrado" {
			return utils.SendNotFoundError(c, "Contato não encontrado")
		}
		return utils.SendInternalServerError(c, "Erro interno do servidor", err.Error())
	}

	return utils.SendSuccessResponse(c, http.StatusOK, "Contato deletado com sucesso", nil)
}
