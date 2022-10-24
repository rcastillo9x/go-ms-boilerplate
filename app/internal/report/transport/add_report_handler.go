package transport

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/rcastillo9x/go-ms-boilerplate/app/internal/report/message/request"
	"github.com/rcastillo9x/go-ms-boilerplate/app/platform/weberror"
	"github.com/rcastillo9x/go-ms-boilerplate/app/platform/weblog"
	"go.uber.org/zap"
)

type AddReportHandler struct {
}

func NewAddReportHandler() AddReportHandler {
	return AddReportHandler{}
}

func (h AddReportHandler) Handle(c *fiber.Ctx) error {
	request := new(request.AddReportRequest)
	weblog.GetInstance().Log(zap.InfoLevel, "Handle Add Report")
	weblog.GetInstance().Log(zap.InfoLevel, "Request", zap.String("body", string(c.Request().Body())))

	errForParser := c.BodyParser(request)
	if errForParser != nil {
		weblog.GetInstance().Log(zap.ErrorLevel, "Request", zap.Error(errForParser))
		weberror.Failure(c, http.StatusBadRequest, weberror.ErrInvalidRequest)
		return nil
	}

	var val = validator.New()
	errForValidation := val.Struct(request)
	if errForValidation != nil {
		weblog.GetInstance().Log(zap.ErrorLevel, "Request", zap.String("Invalid", "atributes"))
		weberror.ErrorForInfraction(c, http.StatusBadRequest, errForValidation)
		return nil
	}

	return c.Status(http.StatusOK).JSON("Add Report Handler")
}
