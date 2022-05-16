package repository

import (
	"gorm.io/gorm"
	"storage/internal/domain/entities"
	"time"
)

type auditRepository struct {
	session *gorm.DB
}

//go:generate mockgen -source=./auditRepository.go -destination=./mocks/auditRepository_mock.go
type AuditRepository interface {
	SaveLog(requestID string, entity string, msg string)
	GetLogs() ([]entities.Audit, error)
}

func NewAuditRepository(session *gorm.DB) AuditRepository {
	return &auditRepository{session: session}
}

func (repository *auditRepository) SaveLog(requestID string, entity string, msg string) {
	_ = repository.session.Create(&entities.Audit{
		RequestID:   requestID,
		ProductID:   entity,
		Description: msg,
		CreatedAt:   time.Now(),
	})
}

func (repository *auditRepository) GetLogs() ([]entities.Audit, error) {
	logs := make([]entities.Audit, 0)
	queryResult := repository.session.Order("created_at desc").Find(&logs)
	return logs, queryResult.Error
}
