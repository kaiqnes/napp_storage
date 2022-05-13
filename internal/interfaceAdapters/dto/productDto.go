package dto

type ProductDto struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Storage   StorageDto `json:"storage"`
	PriceFrom float64    `json:"price_from"`
	PriceTo   float64    `json:"price_to"`
}

type StorageDto struct {
	Total int `json:"total"`
	Corte int `json:"corte"`
}
