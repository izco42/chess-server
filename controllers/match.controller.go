package controllers

import (
	"chess-server/models"
	"chess-server/services"
	"log"
	"github.com/gofiber/fiber/v2"
)

func CreateMatch(c *fiber.Ctx) error {
	match := new(models.Match)
	if err := c.BodyParser(match); err != nil {
		log.Println("Error al leer el cuerpo de la petición:", err)
		return c.Status(400).SendString("Error al leer el cuerpo de la petición")
	}

	matchResponse, err := services.CreateMatch(*match)
	if err != nil {
		return c.Status(500).SendString("Error al crear el partido")
	}
	return c.Status(200).JSON(matchResponse)
}

func GetAllMatches(c *fiber.Ctx) error {  
	matches := services.GetAllMatches()
	return c.Status(200).JSON(matches)
}

func DeleteMatch(c *fiber.Ctx) error{
	var data struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Error al leer el cuerpo de la petición")
	}

	err := services.DeleteMatch(data.ID)
	if err != nil {
		return c.Status(500).SendString("Error al eliminar el partido")
	}
	return c.Status(200).SendString("Partido eliminado")
}


func GetMatchesByPlayerId(c *fiber.Ctx) error {
	var data struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Error al leer el cuerpo de la petición")
	}

	matches,err := services.GetMatchesByPlayerId(data.ID)
	if err != nil {
		return c.Status(500).SendString("Error al obtener los partidos")
	}	
	return c.Status(200).JSON(matches)
}

