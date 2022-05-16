package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	mock_traceability "storage/internal/frameworks/traceability/mocks"
	"storage/internal/interfaceAdapters/presenters"
	mock_useCases "storage/internal/useCases/mocks"
	"testing"
)

func TestProductController_CreateProduct(t *testing.T) {
	scenarios := []productScenario{
		*createProductSuccessfully(),
		*receivesDuplicityError(),
		*receivesRepositoryError(),
		*receivesErrorFromBodyWithEmptyCode(),
		*receivesErrorFromBodyWithEmptyName(),
		*receivesErrorFromBodyWithNegativeTotal(),
		*receivesErrorFromBodyWithNegativeCorte(),
		*receivesErrorFromBodyWithZeroPriceFromAndHigherPriceTo(),
		*receivesErrorFromBodyWithNegativePriceFromAndHigherPriceTo(),
		*receivesErrorFromBodyWithNegativePriceTo(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			router := gin.Default()
			group := router.Group(v1)

			errorPresenter := presenters.NewErrorPresenter()
			productPresenter := presenters.NewProductPresenter(errorPresenter)
			mockLogg := mock_traceability.NewMockApiLogger(ctrl)
			mockUseCase := mock_useCases.NewMockProductUseCase(ctrl)

			testController := NewProductController(group, productPresenter, mockUseCase, mockLogg)
			testController.SetupEndpoints()

			mockLogg.EXPECT().Info(gomock.Any(), gomock.Any())
			if scenario.ShouldMockLoggErr {
				mockLogg.EXPECT().Error(gomock.Any(), gomock.Any())
			}
			if scenario.ShouldMockUseCase {
				mockUseCase.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(scenario.MockUseCaseResult, scenario.MockUseCaseError)
			}

			response := httptest.NewRecorder()
			executeRequest(response, scenario.ReqMethod, getFullUrl(scenario.Uri, ""), scenario.ReqBody, router)

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}

// Para propósito de exemplificação, somente os testes criação de produto.
