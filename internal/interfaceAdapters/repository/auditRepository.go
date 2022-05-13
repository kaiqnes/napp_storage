package repository

import "gorm.io/gorm"

type auditRepository struct {
	session *gorm.DB
}

type AuditRepository interface {
}

func NewAuditRepository(session *gorm.DB) AuditRepository {
	return &auditRepository{session: session}
}
