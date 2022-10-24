package transport

type ReportCommands struct {
	AddReport AddReportHandler
}

type ReportQueries struct {
}

type ReportApplication struct {
	Commands ReportCommands
	Queries  ReportQueries
}
