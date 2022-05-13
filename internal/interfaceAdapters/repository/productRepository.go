package repository

import "gorm.io/gorm"

type productRepository struct {
	session *gorm.DB
}

type ProductRepository interface {
}

func NewProductRepository(session *gorm.DB) ProductRepository {
	return &productRepository{session: session}
}
