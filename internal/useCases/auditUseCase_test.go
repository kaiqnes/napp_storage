package useCases

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	mock_traceability "storage/internal/frameworks/traceability/mocks"
	mock_repository "storage/internal/interfaceAdapters/repository/mocks"
	"testing"
)

func TestAuditUseCase_GetLogs(t *testing.T) {
	scenarios := []auditScenario{
		*receivesOneAuditLog(),
		*receivesTwoAuditLogs(),
		*receivesErrorFromDB(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRepo := mock_repository.NewMockAuditRepository(ctrl)
			mockLogg := mock_traceability.NewMockApiLogger(ctrl)
			testContext, _ := gin.CreateTestContext(httptest.NewRecorder())

			mockLogg.EXPECT().Info(gomock.Any(), gomock.Any())
			mockRepo.EXPECT().GetLogs().Return(scenario.MockRepositoryResponse, scenario.MockRepositoryError)

			//only mock this call if we have an error in this scenario
			if scenario.MockRepositoryError != nil {
				mockLogg.EXPECT().Error(gomock.Any(), gomock.Any())
			}

			useCase := NewAuditUseCase(mockRepo, mockLogg)

			result, errx := useCase.GetLogs(testContext)

			assert.Equal(t, errx, scenario.ExpectError)
			assert.Equal(t, result, scenario.ExpectResult)
		})
	}
}
