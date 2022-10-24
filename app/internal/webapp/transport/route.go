package transport

import "github.com/gofiber/fiber/v2"

func Register(f *fiber.App) {
	h := CheckHandler{}

	f.Get("/status", h.GetStatus)
}
