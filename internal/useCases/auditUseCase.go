package useCases

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"storage/internal/domain/entities"
	"storage/internal/frameworks/errorx"
	"storage/internal/frameworks/traceability"
	"storage/internal/interfaceAdapters/repository"
)

type auditUseCase struct {
	auditRepository repository.AuditRepository
}

type AuditUseCase interface {
	GetLogs(ctx *gin.Context) ([]entities.Audit, errorx.Errorx)
}

func NewAuditUseCase(repository repository.AuditRepository) AuditUseCase {
	return &auditUseCase{auditRepository: repository}
}

func (useCase *auditUseCase) GetLogs(ctx *gin.Context) ([]entities.Audit, errorx.Errorx) {
	traceability.Info(ctx, "Incoming in UseCase GetLogs")

	logs, err := useCase.auditRepository.GetLogs()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve products from DB. err -> %s", err.Error())
		traceability.Error(ctx, errMsg)
		return logs, errorx.NewErrorx(http.StatusInternalServerError, errors.New(errMsg))
	}

	return logs, nil
}
