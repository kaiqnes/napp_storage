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

// healthCheck 	 godoc
// @Summary      This summary endpoint is a health check
// @Description  This description endpoint is a health check
// @Tags         HealthCheck
// @Accept       json
// @Produce      json
// @Success      200 {object} interface{}
// @Router       /health [get]
func (h *healthCheckController) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "https://www.youtube.com/watch?v=xos2MnVxe-c",
	})
}
