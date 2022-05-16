package entities

import (
	"time"
)

type Product struct {
	Code          string  `gorm:"primary_key;size:100"`
	Name          string  `gorm:"size:255"`
	TotalStorage  uint    // Estoque total
	CorteStorage  uint    // Estoque de corte
	OriginalPrice float64 // Preço de
	SalePrice     float64 // Preço por
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
