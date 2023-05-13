package enrrolment

import (
	"github.com/howardhr/Go-Microservices/internal/domain"
	"gorm.io/gorm"
	"log"
)

type (
	Repository interface {
		Create(enroll *domain.Enrrolmet) error
	}
	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(enroll *domain.Enrrolmet) error {
	if err := r.db.Create(enroll).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}
	r.log.Printf("enrolment creado con id:", enroll.ID)
	return nil
}
