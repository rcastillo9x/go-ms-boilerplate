package webapp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	report "github.com/rcastillo9x/go-ms-boilerplate/app/internal/report/transport"
	webapp "github.com/rcastillo9x/go-ms-boilerplate/app/internal/webapp/transport"
)

func New() *fiber.App {
	app := fiber.New()
	configApp(app)
	RegisterURI(app)

	return app
}

func configApp(app *fiber.App) {

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
}

/*
Init handles
*/
func RegisterURI(app *fiber.App) {
	webapp.Register(app)
	report.Register(app)
}
