package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthCheckController struct {
	routes *gin.Engine
}

type HealthCheckController interface {
	SetupEndpoints()
	healthCheck(ctx *gin.Context)
}

func NewHealthCheckController(routes *gin.Engine) HealthCheckController {
	return &healthCheckController{routes: routes}
}

func (h *healthCheckController) SetupEndpoints() {
	h.routes.GET("/health", h.healthCheck)
}

func (h *healthCheckController) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "https://www.youtube.com/watch?v=xos2MnVxe-c",
	})
}
