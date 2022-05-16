package useCases

import (
	"errors"
	"net/http"
	"storage/internal/domain/entities"
	"storage/internal/frameworks/errorx"
	"time"
)

var (
	mockData = time.Date(2000, time.December, 31, 23, 59, 59, 999, time.Local)
)

type auditScenario struct {
	TestName               string
	MockRepositoryResponse []entities.Audit
	MockRepositoryError    error
	ExpectResult           []entities.Audit
	ExpectError            errorx.Errorx
}

func receivesOneAuditLog() *auditScenario {
	log := entities.Audit{
		RequestID:   "abc-123",
		ProductID:   "abc1",
		Description: "product abc1 created successfully",
		CreatedAt:   mockData,
	}
	return &auditScenario{
		TestName:               "AuditUseCase receives a single audit log",
		MockRepositoryResponse: []entities.Audit{log},
		MockRepositoryError:    nil,
		ExpectResult:           []entities.Audit{log},
		ExpectError:            nil,
	}
}

func receivesTwoAuditLogs() *auditScenario {
	log1 := entities.Audit{
		RequestID:   "abc-123",
		ProductID:   "abc1",
		Description: "product abc1 created successfully",
		CreatedAt:   mockData,
	}
	log2 := entities.Audit{
		RequestID:   "abc-456",
		ProductID:   "abc2",
		Description: "product abc2 created successfully",
		CreatedAt:   mockData,
	}

	return &auditScenario{
		TestName:               "AuditUseCase receives two audit logs",
		MockRepositoryResponse: []entities.Audit{log1, log2},
		MockRepositoryError:    nil,
		ExpectResult:           []entities.Audit{log1, log2},
		ExpectError:            nil,
	}
}

func receivesErrorFromDB() *auditScenario {
	dbErr := errors.New("some-db-error")
	errxErr := errors.New("Failed to retrieve products from DB. err -> some-db-error")
	errx := errorx.NewErrorx(http.StatusInternalServerError, errxErr)
	return &auditScenario{
		TestName:               "AuditUseCase receives an error from db layer",
		MockRepositoryResponse: []entities.Audit{},
		MockRepositoryError:    dbErr,
		ExpectResult:           []entities.Audit{},
		ExpectError:            errx,
	}
}
