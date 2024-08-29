package routes
import (
	"chess-server/controllers"
	"github.com/gofiber/fiber/v2"
)

func MatchRoutes(app *fiber.App) {
	match := app.Group("/match")
	match.Get("/r-matchs", controllers.GetAllMatches)
	match.Post("/c-match",controllers.CreateMatch)
	match.Post("/r-match", controllers.GetMatchesByPlayerId)
	match.Delete("/d-match", controllers.DeleteMatch)
}


// {
//     "winnerId": 1,
//     "player2Id": 2,
//     "player3Id": 3,
//     "player4Id": 4,
//     "time": "10:00"
// }

// {
//     "winnerId": 4,
//     "player2Id": 1,
//     "player3Id": 3,
//     "player4Id": 2,
//     "time": "00:03"
// }

