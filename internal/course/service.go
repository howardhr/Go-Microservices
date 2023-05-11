package course

import (
	"log"
	"time"
)

type (
	Service interface {
		Create(Name, StartDate, EndDate string) (*Course, error)
		GetAll(filters Filters, offset, limit int) ([]Course, error)
		Get(id string) (*Course, error)
		Update(id string, name *string, startDate, endDate *string) error
		Delete(id string) error
		Count(filters Filters) (int, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
	Filters struct {
		FirstName string
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(name, startDate, endDate string) (*Course, error) {

	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	course := &Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	if err := s.repo.Create(course); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return course, nil
}
func (s service) GetAll(filters Filters, offset, limit int) ([]Course, error) {
	users, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) Get(id string) (*Course, error) {
	course, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}

func (s service) Update(id string, name *string, startDate, endDate *string) error {
	return s.repo.Update(id, name, startDate, endDate)

}
func (s service) Delete(id string) error {
	return s.repo.Delete(id)

}
