package controllers

import (
  "chess-server/treebasedmodel"
	"github.com/gofiber/fiber/v2"
)


func GetBestMove(c *fiber.Ctx) error {
  var gameState treebasedmodel.GameState
  if err := c.BodyParser(&gameState); err != nil {
    return c.Status(400).SendString("Error al leer el cuerpo de la petición")
  }

  move, err := treebasedmodel.GetBestMove(gameState)
  if err != nil {
    return c.Status(500).SendString("Error al obtener el mejor movimiento")
  }

  return c.Status(200).JSON(move)
}
