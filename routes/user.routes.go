package routes
import (
	"chess-server/controllers"
	"github.com/gofiber/fiber/v2"
)
func UserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/r-users", controllers.GetAllUsers)
	user.Post("/c-user",controllers.CreateUser)
	user.Post("/r-user", controllers.ReadUser)
	user.Delete("/d-user", controllers.DeleteUser)
}
// {
//     "username": "grillo",
//     "password": "123456",
//     "email": "grillo@alumnos.udg.mx"
// }

// {
//     "username": "grillo2",
//     "password" : "123456",
//     "email": "grillo2@alumnos.udg.mx"
// }

// {
//     "username": "grillo3",
//     "password": "123456",
//     "email": "grillo3@alumnos.udg.mx"
// }

// {
//     "username": "grillo4",
//     "password": "123456",
//     "email": "grillo4@alumnos.udg.mx"
// }
