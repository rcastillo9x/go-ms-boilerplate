package commands

import "github.com/rcastillo9x/go-ms-boilerplate/app/internal/report/message/request"

type (
	CreateReportRequestCommand interface {
		Create(command request.AddReportRequest) error
	}
	createReportRequestCommand struct {
	}
)
