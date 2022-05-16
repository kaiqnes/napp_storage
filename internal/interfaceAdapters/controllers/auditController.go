package controllers

import (
	"github.com/gin-gonic/gin"
	"storage/internal/frameworks/traceability"
	"storage/internal/interfaceAdapters/presenters"
	"storage/internal/useCases"
)

type auditController struct {
	routes    *gin.RouterGroup
	presenter presenters.AuditPresenter
	useCase   useCases.AuditUseCase
}

type AuditController interface {
	SetupEndpoints()
	getLogs(ctx *gin.Context)
}

func NewAuditController(routes *gin.RouterGroup, presenter presenters.AuditPresenter, useCase useCases.AuditUseCase) AuditController {
	return &auditController{
		routes:    routes,
		presenter: presenter,
		useCase:   useCase,
	}
}

func (controller *auditController) SetupEndpoints() {
	controller.routes.GET("/logs", controller.getLogs)
}

// getLogs 	 godoc
// @Summary      This endpoint list all logs
// @Description  This endpoint list all logs in details
// @Tags         Audit
// @Accept       json
// @Produce      json
// @Success      200 {object} []dto.AuditDto
// @Failure      500 {object} dto.ErrorOutputDto
// @Router       /api/v1/logs [get]
func (controller *auditController) getLogs(ctx *gin.Context) {
	// Generate RequestID to track logs
	traceability.ValidateRequestID(ctx)
	traceability.Info(ctx, "Incoming in Controller getLogs")

	logs, errx := controller.useCase.GetLogs(ctx)
	if errx != nil {
		controller.presenter.PresentError(ctx, errx.GetError(), errx.GetStatusCode())
		return
	}

	controller.presenter.PresentLogs(ctx, logs)
}
