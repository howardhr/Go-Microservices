package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Enrrolmet struct {
	ID        string     `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID    string     `json:"user_id,omitempty" gorm:"type:char(50);not null"`
	User      *User      `gorm:"user,omitempty"`
	CourseID  string     `json:"course_id,omitempty" gorm:"type:char(50);not null"`
	Course    *Course    `gorm:"user,omitempty"`
	Status    string     `json:"status" gorm:"type:char(50)"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (c *Enrrolmet) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
