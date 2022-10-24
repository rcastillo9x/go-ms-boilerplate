package inmemory

import (
	"errors"
	"sync"

	"github.com/rcastillo9x/go-ms-boilerplate/app/internal/report/model"
)

type ReportInMemory struct {
	reports map[int64]*model.Report
	mutex   *sync.Mutex
}

var (
	Err_ReportNotFound = errors.New("report not found")
)

func NewReportInMemory() *ReportInMemory {
	return &ReportInMemory{
		reports: make(map[int64]*model.Report),
		mutex:   new(sync.Mutex),
	}

}

func (s *ReportInMemory) Create(report *model.Report) (*model.Report, error) {
	s.mutex.Lock()
	s.reports[report.ID] = report
	s.mutex.Unlock()
	return s.reports[report.ID], nil
}

func (s *ReportInMemory) FindByID(ID int64) (*model.Report, error) {
	s.mutex.Lock()
	for _, r := range s.reports {
		if r.ID == ID {
			return s.reports[r.ID], nil
		}
	}
	s.mutex.Unlock()
	return nil, Err_ReportNotFound
}

func (s *ReportInMemory) UpdateByID(report *model.Report) (*model.Report, error) {

	rOriginal, err := s.FindByID(report.ID)

	if err != nil {
		return nil, err
	}
	s.mutex.Lock()
	rOriginal.Definition = report.Definition
	s.reports[report.ID] = rOriginal
	s.mutex.Unlock()

	return rOriginal, nil

}

func (s *ReportInMemory) DeleteByID(id int64) error {
	s.mutex.Lock()
	_, exists := s.reports[id]
	s.mutex.Unlock()
	if !exists {
		return Err_ReportNotFound
	}
	delete(s.reports, id)
	return nil
}

func (s *ReportInMemory) GetAll() (map[int64]*model.Report, error) {
	return s.reports, nil
}
