package controllers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type swaggerController struct {
	router *gin.Engine
}

type SwaggerController interface {
	SetupEndpoints()
}

func NewSwaggerController(routes *gin.Engine) SwaggerController {
	return &swaggerController{
		router: routes,
	}
}

func (s *swaggerController) SetupEndpoints() {
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
