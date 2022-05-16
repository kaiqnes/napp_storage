package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckController(t *testing.T) {
	scenarios := []healthCheckScenario{
		*getHealthCheckUP(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			router := gin.Default()

			testController := NewHealthCheckController(router)
			testController.SetupEndpoints()

			response := httptest.NewRecorder()
			executeRequest(response, http.MethodGet, getFullUrl(scenario.Uri, ""), emptyBody, router)

			assert.Equal(t, response.Body.String(), scenario.ExpectResponse)
			assert.Equal(t, response.Result().StatusCode, scenario.ExpectStatus)
		})
	}
}
