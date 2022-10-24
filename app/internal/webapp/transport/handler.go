package transport

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CheckHandler struct {
}

func (ch *CheckHandler) GetStatus(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("UP")

}
