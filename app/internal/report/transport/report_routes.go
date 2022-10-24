package transport

import (
	"github.com/gofiber/fiber/v2"
)

func Register(f *fiber.App) {
	ra := ReportApplication{}

	f.Post("/api/v1/reports", ra.Commands.AddReport.Handle)

}
