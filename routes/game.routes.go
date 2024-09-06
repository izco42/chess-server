package routes
import (
  "github.com/gofiber/fiber/v2"
  "chess-server/controllers"
)

func GameRoutes(app *fiber.App) { 
  game := app.Group("/game")
  game.Post("/get-best-move", controllers.GetBestMove)
}
