package request

type AddReportRequest struct {
	Definition string `json:"definition" validate:"required"`
}
