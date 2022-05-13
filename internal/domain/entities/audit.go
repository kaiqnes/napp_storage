package entities

import "gorm.io/gorm"

type Audit struct {
	gorm.Model
	ProductID   uint
	Description string
}
