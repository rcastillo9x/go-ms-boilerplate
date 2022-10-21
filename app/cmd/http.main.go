package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rcastillo9x/go-ms-boilerplate/app/internal/webapp"
	"github.com/rcastillo9x/go-ms-boilerplate/app/platform/enviroment"
	"github.com/rcastillo9x/go-ms-boilerplate/app/platform/weblog"
)

func main() {

	wl := weblog.GetInstance()
	wl.Info("Inicializando webapp")

	// Initialize environment variables
	HTTP_PORT := enviroment.GetHTTPPort("HTTP_PORT")

	// Initialize Http server
	app := webapp.New()

	// Start server on port 3000
	app.Listen(HTTP_PORT)

}

func registerURI(app *fiber.App, db *sqlx.DB) {
	webapp.Register(app)

}
