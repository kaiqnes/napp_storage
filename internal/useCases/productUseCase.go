package useCases

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/internal/domain/entities"
	"storage/internal/frameworks/errorx"
	"storage/internal/frameworks/traceability"
	"storage/internal/interfaceAdapters/dto"
	"storage/internal/interfaceAdapters/repository"
	"strings"
	"time"
)

type productUseCase struct {
	productRepository repository.ProductRepository
	auditRepository   repository.AuditRepository
}

type ProductUseCase interface {
	GetProducts(ctx *gin.Context, filterParam string, limit, offset int) ([]entities.Product, errorx.Errorx)
	GetProduct(ctx *gin.Context, code string) (entities.Product, errorx.Errorx)
	CreateProduct(ctx *gin.Context, productDto dto.ProductInputDto) (entities.Product, errorx.Errorx)
	UpdateProduct(ctx *gin.Context, code string, productDto dto.ProductInputDto) (entities.Product, errorx.Errorx)
	DeleteProduct(ctx *gin.Context, code string) errorx.Errorx
}

func NewProductUseCase(productRepository repository.ProductRepository, auditRepository repository.AuditRepository) ProductUseCase {
	return &productUseCase{
		productRepository: productRepository,
		auditRepository:   auditRepository,
	}
}

func (useCase *productUseCase) GetProducts(ctx *gin.Context, filterParam string, limit, offset int) ([]entities.Product, errorx.Errorx) {
	traceability.Info(ctx, "Incoming in UseCase GetProducts")

	products, err := useCase.productRepository.GetProducts(filterParam, limit, offset)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve products from DB. err -> %s", err.Error())
		traceability.Error(ctx, errMsg)
		return products, errorx.NewErrorx(http.StatusInternalServerError, errors.New(errMsg))
	}

	return products, nil
}

func (useCase *productUseCase) GetProduct(ctx *gin.Context, code string) (entities.Product, errorx.Errorx) {
	traceability.Info(ctx, fmt.Sprintf("Incoming in UseCase GetProduct with code %s", code))

	product, err := useCase.productRepository.GetProduct(code)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve products from DB. err -> %s", err.Error())
		traceability.Error(ctx, errMsg)
		return product, errorx.NewErrorx(http.StatusInternalServerError, errors.New(errMsg))
	}

	if product.Code != code {
		errMsg := fmt.Sprintf("Failed to retrieve product with code %s", code)
		traceability.Error(ctx, errMsg)
		return product, errorx.NewErrorx(http.StatusNotFound, errors.New(errMsg))
	}

	return product, nil
}

func (useCase *productUseCase) CreateProduct(ctx *gin.Context, productDto dto.ProductInputDto) (createdProduct entities.Product, errx errorx.Errorx) {
	var (
		productEntity = entities.Product{
			Code:          productDto.Code,
			Name:          productDto.Name,
			TotalStorage:  productDto.Storage.Total,
			CorteStorage:  productDto.Storage.Corte,
			OriginalPrice: productDto.PriceFrom,
			SalePrice:     productDto.PriceTo,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		err error
	)

	traceability.Info(ctx, "Incoming in UseCase CreateProduct")
	createdProduct, err = useCase.productRepository.CreateProduct(productEntity)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			errMsg := fmt.Sprintf("error 1062: Duplicate entry '%s' for key 'code'", productEntity.Code)
			traceability.Error(ctx, errMsg)
			errx = errorx.NewErrorx(http.StatusBadRequest, errors.New(errMsg))
		} else {
			traceability.Error(ctx, err.Error())
			errx = errorx.NewErrorx(http.StatusInternalServerError, err)
		}
		return
	}

	useCase.audit(ctx, productEntity.Code, "product created successfully")
	return
}

func (useCase *productUseCase) UpdateProduct(ctx *gin.Context, code string, productDto dto.ProductInputDto) (productEntity entities.Product, errx errorx.Errorx) {
	productEntity = entities.Product{
		Code:          productDto.Code,
		Name:          productDto.Name,
		TotalStorage:  productDto.Storage.Total,
		CorteStorage:  productDto.Storage.Corte,
		OriginalPrice: productDto.PriceFrom,
		SalePrice:     productDto.PriceTo,
		UpdatedAt:     time.Now(),
	}

	traceability.Info(ctx, "Incoming in UseCase UpdateProduct")

	if err := useCase.productRepository.UpdateProduct(code, productEntity); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			errMsg := fmt.Sprintf("error 1062: Duplicate entry '%s' for key 'code'", productEntity.Code)
			traceability.Error(ctx, errMsg)
			errx = errorx.NewErrorx(http.StatusBadRequest, errors.New(errMsg))
		} else {
			traceability.Error(ctx, err.Error())
			errx = errorx.NewErrorx(http.StatusInternalServerError, err)
		}
		return
	}

	useCase.audit(ctx, productEntity.Code, "product update successfully")
	return
}

func (useCase *productUseCase) DeleteProduct(ctx *gin.Context, code string) (errx errorx.Errorx) {
	traceability.Info(ctx, "Incoming in UseCase DeleteProduct")

	if rows, err := useCase.productRepository.DeleteProduct(code); err != nil {
		traceability.Error(ctx, err.Error())
		return errorx.NewErrorx(http.StatusInternalServerError, err)
	} else if rows == 0 {
		errMsg := fmt.Sprintf("product with code %s was not found in DB.", code)
		traceability.Error(ctx, errMsg)
		return errorx.NewErrorx(http.StatusNotFound, errors.New(errMsg))
	}

	useCase.audit(ctx, code, "product deleted successfully")
	return
}

func (useCase *productUseCase) audit(ctx *gin.Context, code string, errMsg string) {
	go useCase.auditRepository.SaveLog(traceability.GetRequestID(ctx), code, errMsg)
}
