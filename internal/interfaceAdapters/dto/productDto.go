package dto

import "time"

type ProductInputDto struct {
	Code      string          `json:"code"`
	Name      string          `json:"name"`
	Storage   StorageInputDto `json:"storage"`
	PriceFrom float64         `json:"price_from"`
	PriceTo   float64         `json:"price_to"`
}

type StorageInputDto struct {
	Total uint `json:"total"`
	Corte uint `json:"corte"`
}

type ProductOutputDto struct {
	Code      string           `json:"code"`
	Name      string           `json:"name"`
	Storage   StorageOutputDto `json:"storage"`
	PriceFrom float64          `json:"price_from"`
	PriceTo   float64          `json:"price_to"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type StorageOutputDto struct {
	Total     uint `json:"total"`
	Corte     uint `json:"corte"`
	Available uint `json:"disponivel"`
}
