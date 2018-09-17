package domain

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model is base model
type Model struct {
	ID        UUID       `sql:",type:uuid" json:"id"`
	CreatedAt time.Time  `sql:"default:now()" json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// BeforeCreate prepare data before create data
func (at *Model) BeforeCreate(scope *gorm.Scope) error {
	emptyID, _ := UUIDFromString("")
	if at.ID == emptyID {
		scope.SetColumn("ID", uuid.NewV4())
	}
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil

}
