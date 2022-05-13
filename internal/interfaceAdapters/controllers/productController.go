package controllers

import (
	"github.com/gin-gonic/gin"
	"storage/internal/interfaceAdapters/presenters"
	"storage/internal/useCases"
)

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
	controller.routes.GET("/products/:productID", controller.getProduct)
	controller.routes.POST("/products", controller.createProduct)
	controller.routes.PUT("/products/:productID", controller.updateProduct) //TODO: PUT OR PATCH?
	controller.routes.DELETE("/products/:productID", controller.deleteProduct)
}

func (controller *productController) getProducts(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *productController) getProduct(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *productController) createProduct(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *productController) updateProduct(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *productController) deleteProduct(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
