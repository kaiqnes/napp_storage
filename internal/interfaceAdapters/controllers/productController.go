package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/internal/domain/entities"
	"storage/internal/frameworks/traceability"
	"storage/internal/interfaceAdapters/dto"
	"storage/internal/interfaceAdapters/presenters"
	"storage/internal/useCases"
	"strconv"
)

const defaultLimit = 10

type productController struct {
	routes    *gin.RouterGroup
	presenter presenters.ProductPresenter
	useCase   useCases.ProductUseCase
}

type ProductController interface {
	SetupEndpoints()
	getProducts(ctx *gin.Context)
	getProduct(ctx *gin.Context)
	createProduct(ctx *gin.Context)
	updateProduct(ctx *gin.Context)
	deleteProduct(ctx *gin.Context)
}

func NewProductController(routes *gin.RouterGroup, presenter presenters.ProductPresenter, useCase useCases.ProductUseCase) ProductController {
	return &productController{
		routes:    routes,
		presenter: presenter,
		useCase:   useCase,
	}
}

func (controller *productController) SetupEndpoints() {
	controller.routes.GET("/products", controller.getProducts)
	controller.routes.GET("/products/:productCode", controller.getProduct)
	controller.routes.POST("/products", controller.createProduct)
	controller.routes.PUT("/products/:productCode", controller.updateProduct)
	controller.routes.DELETE("/products/:productCode", controller.deleteProduct)
}

func (controller *productController) getProducts(ctx *gin.Context) {
	// Generate RequestID to track logs
	traceability.ValidateRequestID(ctx)
	traceability.Info(ctx, "Incoming in Controller getProducts")

	filterParam, limit, offset, err := getParamsToPaginateAndFilter(ctx)
	if err != nil {
		controller.presenter.PresentError(ctx, err, http.StatusBadRequest)
		return
	}

	products, errx := controller.useCase.GetProducts(ctx, filterParam, limit, offset)
	if errx != nil {
		controller.presenter.PresentError(ctx, errx.GetError(), errx.GetStatusCode())
		return
	}

	controller.presenter.PresentGetProducts(ctx, products)
}

func (controller *productController) getProduct(ctx *gin.Context) {
	// Generate RequestID to track logs
	traceability.ValidateRequestID(ctx)
	traceability.Info(ctx, "from controller in method getProduct")

	productCode := ctx.Param("productCode")

	product, errx := controller.useCase.GetProduct(ctx, productCode)
	if errx != nil {
		controller.presenter.PresentError(ctx, errx.GetError(), errx.GetStatusCode())
		return
	}

	controller.presenter.PresentGetProducts(ctx, []entities.Product{product})
}

func (controller *productController) createProduct(ctx *gin.Context) {
	// Generate RequestID to track logs
	traceability.ValidateRequestID(ctx)
	traceability.Info(ctx, "from controller in method createProduct")

	productDto, errDto, err := isValidRequestBody(ctx)
	if err != nil {
		controller.presenter.PresentError(ctx, err, http.StatusBadRequest)
		return
	}
	if len(errDto.Fields) > 0 {
		controller.presenter.PresentErrors(ctx, errDto, http.StatusBadRequest)
		return
	}

	createdProduct, errx := controller.useCase.CreateProduct(ctx, productDto)
	if errx != nil {
		controller.presenter.PresentError(ctx, errx.GetError(), errx.GetStatusCode())
		return
	}

	controller.presenter.PresentCreateProduct(ctx, createdProduct)
}

func (controller *productController) updateProduct(ctx *gin.Context) {
	// Generate RequestID to track logs
	traceability.ValidateRequestID(ctx)
	traceability.Info(ctx, "from controller in method updateProduct")

	productCode := ctx.Param("productCode")

	productDto, errDto, err := isValidRequestBody(ctx)
	if err != nil {
		controller.presenter.PresentError(ctx, err, http.StatusBadRequest)
		return
	}
	if len(errDto.Fields) > 0 {
		controller.presenter.PresentErrors(ctx, errDto, http.StatusBadRequest)
		return
	}

	createdProduct, errx := controller.useCase.UpdateProduct(ctx, productCode, productDto)
	if errx != nil {
		controller.presenter.PresentError(ctx, errx.GetError(), errx.GetStatusCode())
		return
	}

	controller.presenter.PresentUpdateProduct(ctx, createdProduct)
}

func (controller *productController) deleteProduct(ctx *gin.Context) {
	// Generate RequestID to track logs
	traceability.ValidateRequestID(ctx)
	traceability.Info(ctx, "from controller in method deleteProduct")

	productCode := ctx.Param("productCode")

	if errx := controller.useCase.DeleteProduct(ctx, productCode); errx != nil {
		controller.presenter.PresentError(ctx, errx.GetError(), errx.GetStatusCode())
		return
	}

	controller.presenter.PresentDeleteProduct(ctx)
}

func isValidRequestBody(ctx *gin.Context) (product dto.ProductInputDto, errDto dto.ErrorsOutputDto, err error) {
	err = ctx.BindJSON(&product)
	if err != nil {
		return
	}

	if product.Code == "" {
		errDto.Fields = append(errDto.Fields, "code must not be empty")
	}
	if product.Name == "" {
		errDto.Fields = append(errDto.Fields, "name must not be empty")
	}
	if product.Storage.Total < 0 {
		errDto.Fields = append(errDto.Fields, "storage.total must not be less than zero")
	}
	if product.Storage.Corte < 0 {
		errDto.Fields = append(errDto.Fields, "storage.corte must not be less than zero")
	}
	if product.PriceFrom < 0 {
		errDto.Fields = append(errDto.Fields, "price_from must not be less than zero")
	}
	if product.PriceTo < 0 {
		errDto.Fields = append(errDto.Fields, "price_to must not be less than zero")
	}
	if product.PriceFrom < product.PriceTo {
		errDto.Fields = append(errDto.Fields, "price_from must not be lower than price_to")
	}

	if len(errDto.Fields) > 0 {
		errDto.Message = "incorrect value(s) received."
	}

	return
}

func getParamsToPaginateAndFilter(ctx *gin.Context) (string, int, int, error) {
	var (
		filterParam     = ctx.Query("q")
		limit           = ctx.Query("limit")
		offset          = ctx.Query("offset")
		iLimit, iOffset int
	)

	if limit != "" {
		parsedLimit, err := strconv.ParseUint(limit, 10, 32)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to parse limit. Received: %s | It must be numeric, integer and positive", limit)
			traceability.Error(ctx, errMsg)
			return filterParam, iLimit, iOffset, errors.New(errMsg)
		}
		iLimit = int(parsedLimit)
	}

	if offset != "" {
		parsedOffset, err := strconv.ParseUint(offset, 10, 32)
		if err != nil {
			errMsg := fmt.Sprintf("Failed to parse offset. Received: %s | It must be numeric, integer and positive", offset)
			traceability.Error(ctx, errMsg)
			return filterParam, iLimit, iOffset, errors.New(errMsg)
		}
		iOffset = int(parsedOffset)
	}

	if iLimit == 0 {
		traceability.Info(ctx, "Using default limit")
		iLimit = defaultLimit
	}

	return filterParam, iLimit, iOffset, nil
}
