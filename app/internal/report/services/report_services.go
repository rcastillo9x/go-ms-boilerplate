package services

type ReportCommandService struct {
	AddReport AddReportHandler
}

type ReportQueriesServices struct {
}

type ReportService struct {
	Commands ReportCommandService
	Queries  ReportQueriesServices
}
