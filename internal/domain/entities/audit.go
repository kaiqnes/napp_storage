package entities

import "time"

type Audit struct {
	RequestID   string `gorm:"primary_key;size:100"`
	ProductID   string
	Description string
	CreatedAt   time.Time
}
