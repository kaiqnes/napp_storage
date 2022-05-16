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
	logg            traceability.ApiLogger
}

type AuditUseCase interface {
	GetLogs(ctx *gin.Context) ([]entities.Audit, errorx.Errorx)
}

func NewAuditUseCase(repository repository.AuditRepository, logger traceability.ApiLogger) AuditUseCase {
	return &auditUseCase{auditRepository: repository, logg: logger}
}

func (useCase *auditUseCase) GetLogs(ctx *gin.Context) ([]entities.Audit, errorx.Errorx) {
	useCase.logg.Info(ctx, "Incoming in UseCase GetLogs")

	logs, err := useCase.auditRepository.GetLogs()
	if err != nil {
		errMsg := fmt.Sprintf("Failed to retrieve products from DB. err -> %s", err.Error())
		useCase.logg.Error(ctx, errMsg)
		return logs, errorx.NewErrorx(http.StatusInternalServerError, errors.New(errMsg))
	}

	return logs, nil
}
