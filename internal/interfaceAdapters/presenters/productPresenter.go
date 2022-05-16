package presenters

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/internal/domain/entities"
	"storage/internal/interfaceAdapters/dto"
)

type ProductPresenter interface {
	PresentError(ctx *gin.Context, err error, statusCode int)
	PresentErrors(ctx *gin.Context, errDto dto.ErrorsOutputDto, statusCode int)
	PresentGetProducts(ctx *gin.Context, products []entities.Product)
	PresentCreateProduct(ctx *gin.Context, product entities.Product)
	PresentUpdateProduct(ctx *gin.Context, product entities.Product)
	PresentDeleteProduct(ctx *gin.Context)
}

type productPresenter struct {
	errorPresenter ErrorPresenter
}

func NewProductPresenter(errorPresenter ErrorPresenter) ProductPresenter {
	return &productPresenter{errorPresenter: errorPresenter}
}

func (presenter *productPresenter) PresentError(ctx *gin.Context, err error, statusCode int) {
	presenter.errorPresenter.PresentError(ctx, err, statusCode)
}

func (presenter *productPresenter) PresentErrors(ctx *gin.Context, errDto dto.ErrorsOutputDto, statusCode int) {
	presenter.errorPresenter.PresentErrors(ctx, errDto, statusCode)
}

func (presenter *productPresenter) PresentGetProducts(ctx *gin.Context, products []entities.Product) {
	productsDto := make([]dto.ProductOutputDto, 0)

	for _, product := range products {
		productDto := dto.ProductOutputDto{
			Code: product.Code,
			Name: product.Name,
			Storage: dto.StorageOutputDto{
				Total:     product.TotalStorage,
				Corte:     product.CorteStorage,
				Available: product.TotalStorage - product.CorteStorage,
			},
			PriceFrom: product.OriginalPrice,
			PriceTo:   product.SalePrice,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}
		productsDto = append(productsDto, productDto)
	}

	ctx.JSON(http.StatusOK, productsDto)
}

func (presenter *productPresenter) PresentCreateProduct(ctx *gin.Context, product entities.Product) {
	productsDto := []dto.ProductOutputDto{{
		Code: product.Code,
		Name: product.Name,
		Storage: dto.StorageOutputDto{
			Total:     product.TotalStorage,
			Corte:     product.CorteStorage,
			Available: product.TotalStorage - product.CorteStorage,
		},
		PriceFrom: product.OriginalPrice,
		PriceTo:   product.SalePrice,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}}

	ctx.JSON(http.StatusCreated, productsDto)
}

func (presenter *productPresenter) PresentUpdateProduct(ctx *gin.Context, product entities.Product) {
	productsDto := []dto.ProductOutputDto{{
		Code: product.Code,
		Name: product.Name,
		Storage: dto.StorageOutputDto{
			Total:     product.TotalStorage,
			Corte:     product.CorteStorage,
			Available: product.TotalStorage - product.CorteStorage,
		},
		PriceFrom: product.OriginalPrice,
		PriceTo:   product.SalePrice,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}}

	ctx.JSON(http.StatusOK, productsDto)
}

func (presenter *productPresenter) PresentDeleteProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}
