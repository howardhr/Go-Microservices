package course

import (
	"fmt"
	"github.com/howardhr/Go-Microservices/internal/domain"
	"gorm.io/gorm"
	"log"
	"strings"
)

type (
	Repository interface {
		Create(course *domain.Course) error
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		Update(id string, name *string, startDate, endDate *string) error
		Delete(id string) error
		Count(filters Filters) (int, error)
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

func (repo *repo) Create(course *domain.Course) error {

	if err := repo.db.Create(course).Error; err != nil {
		fmt.Println(err)
		return err
	}
	repo.log.Println("Curso creado con id:", course.ID)
	return nil
}

func (repo *repo) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	var c []domain.Course

	tx := repo.db.Model(&c)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(name) like ?", filters.FirstName)
	}

	return tx
}

func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (repo *repo) Get(id string) (*domain.Course, error) {
	user := domain.Course{ID: id}
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) Update(id string, name *string, startDate, endDate *string) error {
	values := make(map[string]interface{})
	if name != nil {
		values["name"] = *name
	}
	if startDate != nil {
		values["start_date"] = *startDate
	}
	if endDate != nil {
		values["end_date"] = *endDate
	}
	if err := r.db.Model(&domain.Course{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) Delete(id string) error {
	user := domain.Course{ID: id}
	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
