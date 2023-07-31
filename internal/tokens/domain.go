package tokens

import (
	"time"

	"gorm.io/gorm"
)

type Tokens struct {
	ID          string         `json:"id"`
	Image       string         `json:"image"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
