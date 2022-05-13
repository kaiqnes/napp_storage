package useCases

import (
	"storage/internal/interfaceAdapters/repository"
)

type productUseCase struct {
	productRepository repository.ProductRepository
	auditRepository   repository.AuditRepository
}

type ProductUseCase interface {
}

func NewProductUseCase(productRepository repository.ProductRepository, auditRepository repository.AuditRepository) ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
		auditRepository:   auditRepository,
	}
}
