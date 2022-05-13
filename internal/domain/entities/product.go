package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code          string
	Name          string
	TotalStorage  uint    // Estoque total
	CorteStorage  uint    // Estoque de corte
	OriginalPrice float64 // Preço de
	SalePrice     float64 // Preço por
}
