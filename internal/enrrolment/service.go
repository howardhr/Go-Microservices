package enrrolment

import (
	"errors"
	"github.com/howardhr/Go-Microservices/internal/course"
	"github.com/howardhr/Go-Microservices/internal/domain"
	"github.com/howardhr/Go-Microservices/internal/user"
	"log"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrrolmet, error)
	}
	service struct {
		log       *log.Logger
		userSrv   user.Service //llamada al servicio
		courseSrv course.Service
		repo      Repository
	}
)

func NewService(l *log.Logger, userSrv user.Service, courseSrv course.Service, repo Repository) Service {
	return &service{
		log:       l,
		userSrv:   userSrv, //se le pasan los valores que vienen por parametro
		courseSrv: courseSrv,
		repo:      repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrrolmet, error) {
	enroll := &domain.Enrrolmet{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}
	if _, err := s.userSrv.Get(enroll.UserID); err != nil {
		return nil, errors.New("el usuario no existe")
	}
	if _, err := s.courseSrv.Get(enroll.CourseID); err != nil { //se realiza el get de ambos
		return nil, errors.New("el curso no existe")
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}
	return enroll, nil
}
