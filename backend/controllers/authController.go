package controllers

import (
	"github.com/gofiber/fiber/v3"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao processar a requisição",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Usuário registrado com sucesso",
	})
}
