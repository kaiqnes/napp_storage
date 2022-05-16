package presenters

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/internal/domain/entities"
	"storage/internal/interfaceAdapters/dto"
)

type AuditPresenter interface {
	PresentError(ctx *gin.Context, err error, statusCode int)
	PresentLogs(ctx *gin.Context, logs []entities.Audit)
}

type auditPresenter struct {
	errorPresenter ErrorPresenter
}

func NewAuditPresenter(errorPresenter ErrorPresenter) AuditPresenter {
	return &auditPresenter{errorPresenter: errorPresenter}
}

func (presenter *auditPresenter) PresentError(ctx *gin.Context, err error, statusCode int) {
	presenter.errorPresenter.PresentError(ctx, err, statusCode)
}

func (presenter *auditPresenter) PresentLogs(ctx *gin.Context, logs []entities.Audit) {
	logsDto := make([]dto.AuditDto, 0)

	for _, log := range logs {
		logsDto = append(logsDto, dto.AuditDto{
			RequestID:   log.RequestID,
			ProductID:   log.ProductID,
			Description: log.Description,
			CreatedAt:   log.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, logsDto)
}
