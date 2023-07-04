package domain

import (
	uuid "github.com/google/uuid"
	"time"
)

type Example struct {
	ID        int        `gorm:"primaryKey"`
	Uuid      uuid.UUID  `json:"uuid" gorm:"not null"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BaseExampleResponse struct {
	Uuid      uuid.UUID  `json:"uuid"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}

type ExampleResponse struct {
	Success   bool                `json:"success"`
	ErrorCode int16               `json:"error_code"`
	Message   string              `json:"message"`
	Data      BaseExampleResponse `json:"data"`
}

type ListExampleResponse struct {
	Success   bool                  `json:"success"`
	ErrorCode int16                 `json:"error_code"`
	Message   string                `json:"message"`
	Data      []BaseExampleResponse `json:"data"`
	Metadata  MetadataResponse      `json:"metadata"`
}

type ExampleQuery struct{}
