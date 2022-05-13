package di

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"storage/internal/interfaceAdapters/controllers"
	"storage/internal/interfaceAdapters/presenters"
	"storage/internal/interfaceAdapters/repository"
	"storage/internal/useCases"
)

type dependencyInjection struct {
	routes *gin.Engine
	db     *gorm.DB
}

func NewDependencyInjection(routes *gin.Engine, session *gorm.DB) *dependencyInjection {
	return &dependencyInjection{
		routes: routes,
		db:     session,
	}
}

func (di *dependencyInjection) SetupDependencies() {
	di.injectStructuralResources()
	di.injectPublicResources()
}

func (di *dependencyInjection) injectPublicResources() {
	publicGroup := di.routes.Group("/api/v1")
	errorPresenter := presenters.NewErrorPresenter()
	auditRepository := repository.NewAuditRepository(di.db)

	/* Product Resource */
	productPresenter := presenters.NewProductPresenter(errorPresenter)
	productRepository := repository.NewProductRepository(di.db)
	productUseCase := useCases.NewProductUseCase(productRepository, auditRepository)
	productController := controllers.NewProductController(publicGroup, productPresenter, productUseCase)
	productController.SetupEndpoints()
}

func (di *dependencyInjection) injectStructuralResources() {
	//HealthCheck
	//Swagger
}
