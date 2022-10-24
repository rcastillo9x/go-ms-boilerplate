package weberror

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var (
	// Global Errors
	ErrInvalidRequest = errors.New("invalid request")
)

type (
	// Error for validation
	ValidationError struct {
		Field string
		Rule  string
		Value string
	}

	// Common reponse
	errorResponse struct {
		Status      int                `json:"status"`
		Code        string             `json:"code"`
		Message     string             `json:"message"`
		Infractions []*ValidationError `json:"infractions"`
	}

	response struct {
		Data interface{} `json:"data"`
	}
)

func Success(c *fiber.Ctx, status int, data interface{}) {
	c.JSON(response{
		Data: data,
	})
}

func Failure(c *fiber.Ctx, status int, err error) *fiber.Ctx {
	e := &errorResponse{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	}

	c.Status(status)
	c.JSON(e)
	return c

}

func ErrorForInfraction(c *fiber.Ctx, status int, err error) *fiber.Ctx {

	var infractions []*ValidationError

	for _, err := range err.(validator.ValidationErrors) {
		var el ValidationError
		el.Field = err.Field()
		el.Rule = err.Tag()
		el.Value = err.Param()
		infractions = append(infractions, &el)
	}

	e := &errorResponse{
		Message:     errors.New("invalid atributes").Error(),
		Status:      status,
		Code:        http.StatusText(status),
		Infractions: infractions,
	}

	c.Status(status)

	c.JSON(e)

	return c

}
