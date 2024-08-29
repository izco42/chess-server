package controllers

import (
	"chess-server/models"
	"chess-server/services"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}

func GetAllUsers(c *fiber.Ctx) error {
	users := services.GetAllUsers()
	return c.Status(200).JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Error al leer el cuerpo de la petición")
	}
	user, err := services.CreateUser(*user)
	if err != nil {
		return c.Status(500).SendString("Error al crear el usuario")
	}
	return c.Status(200).JSON(user)
}

func ReadUser(c *fiber.Ctx) error {
	userLogin := new(models.UserLogin)
	if err := c.BodyParser(userLogin); err != nil {
		return c.Status(400).SendString("Error al leer el cuerpo de la petición")
	}
	user, err := services.LoginUser(*userLogin)
	if err != nil {
		return c.Status(401).SendString("Usuario o contraseña incorrectos")
	}
	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var data struct {
		ID uint `json:"id"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Error al leer el cuerpo de la petición")
	}
	err := services.DeleteUser(data.ID)
	if err != nil {
		return c.Status(500).SendString("Error al eliminar el usuario")
	}
	return c.Status(200).SendString("Usuario eliminado")
}

