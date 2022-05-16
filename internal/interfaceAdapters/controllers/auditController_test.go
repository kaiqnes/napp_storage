package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	mock_traceability "storage/internal/frameworks/traceability/mocks"
	"storage/internal/interfaceAdapters/presenters"
	mock_useCases "storage/internal/useCases/mocks"
	"testing"
)

func TestAuditController_GetLogs(t *testing.T) {
	scenarios := []auditScenario{
		*getOneAuditLog(),
		*getTwoAuditLogs(),
		*getErrorFromUseCase(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			router := gin.Default()
			group := router.Group(v1)

			errorPresenter := presenters.NewErrorPresenter()
			auditPresenter := presenters.NewAuditPresenter(errorPresenter)
			mockLogg := mock_traceability.NewMockApiLogger(ctrl)
			mockUseCase := mock_useCases.NewMockAuditUseCase(ctrl)

			testController := NewAuditController(group, auditPresenter, mockUseCase, mockLogg)
			testController.SetupEndpoints()

			mockLogg.EXPECT().Info(gomock.Any(), gomock.Any())
			mockUseCase.EXPECT().GetLogs(gomock.Any()).Return(scenario.MockUseCaseLogs, scenario.MockUseCaseError)

			response := httptest.NewRecorder()
			executeRequest(response, http.MethodGet, getFullUrl(scenario.Uri, ""), emptyBody, router)

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}
