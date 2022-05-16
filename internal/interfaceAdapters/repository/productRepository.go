package repository

import (
	"fmt"
	"gorm.io/gorm"
	"storage/internal/domain/entities"
)

type productRepository struct {
	session *gorm.DB
}

type ProductRepository interface {
	GetProducts(filterParam string, limit, offset int) ([]entities.Product, error)
	GetProduct(code string) (entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(code string, product entities.Product) error
	DeleteProduct(code string) (int64, error)
}

func NewProductRepository(session *gorm.DB) ProductRepository {
	return &productRepository{session: session}
}

func (repository *productRepository) GetProducts(filterParam string, limit, offset int) ([]entities.Product, error) {
	products := make([]entities.Product, 0)
	query := repository.session.Limit(limit).Offset(offset)

	if filterParam != "" {
		likeableParam := fmt.Sprintf("%%%s%%", filterParam)
		query.Or("code LIKE ?", likeableParam)
		query.Or("name LIKE ?", likeableParam)
	}

	queryResult := query.Order("updated_at desc").Find(&products)

	return products, queryResult.Error
}

func (repository *productRepository) GetProduct(code string) (entities.Product, error) {
	var product entities.Product
	queryResult := repository.session.Where("code", code).Find(&product)
	return product, queryResult.Error
}

func (repository *productRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	queryResult := repository.session.Create(&product)
	return product, queryResult.Error
}

func (repository *productRepository) UpdateProduct(code string, product entities.Product) error {
	queryResult := repository.session.Model(entities.Product{}).Where("code = ?", code).Updates(product)
	return queryResult.Error
}

func (repository *productRepository) DeleteProduct(code string) (int64, error) {
	queryResult := repository.session.Delete(entities.Product{Code: code})
	return queryResult.RowsAffected, queryResult.Error
}
