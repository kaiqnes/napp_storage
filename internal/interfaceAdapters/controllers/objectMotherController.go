package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"storage/internal/domain/entities"
	"storage/internal/frameworks/errorx"
	"time"
)

const (
	baseUrl        = "http://localhost:8080"
	v1             = "/api/v1"
	productsV1Uri  = v1 + "/products"
	logsV1Uri      = v1 + "/logs"
	healthCheckUri = "/health"
	emptyBody      = ""
)

var (
	mockData = time.Date(2000, time.December, 31, 23, 59, 59, 999, time.Local)
)

type healthCheckScenario struct {
	TestName       string
	Uri            string
	ExpectStatus   int
	ExpectResponse string
}

type auditScenario struct {
	TestName         string
	Uri              string
	MockUseCaseLogs  []entities.Audit
	MockUseCaseError errorx.Errorx
	ExpectStatus     int
	ExpectResponse   string
}

type productScenario struct {
	TestName          string
	Uri               string
	ReqBody           string
	ReqMethod         string
	ShouldMockUseCase bool
	ShouldMockLoggErr bool
	MockUseCaseResult entities.Product
	MockUseCaseError  errorx.Errorx
	ExpectStatus      int
	ExpectResponse    string
}

func getHealthCheckUP() *healthCheckScenario {
	return &healthCheckScenario{
		TestName:       "Get Health Check result",
		Uri:            healthCheckUri,
		ExpectStatus:   http.StatusOK,
		ExpectResponse: `{"response":"https://www.youtube.com/watch?v=xos2MnVxe-c"}`,
	}
}

func getOneAuditLog() *auditScenario {
	log := entities.Audit{
		RequestID:   "abc-123",
		ProductID:   "abc1",
		Description: "product abc1 created successfully",
		CreatedAt:   mockData,
	}

	useCaseLogs := []entities.Audit{log}

	return &auditScenario{
		TestName:         "Get one product log from Audit endpoint",
		Uri:              logsV1Uri,
		MockUseCaseLogs:  useCaseLogs,
		MockUseCaseError: nil,
		ExpectStatus:     http.StatusOK,
		ExpectResponse:   `[{"request_id":"abc-123","product_id":"abc1","description":"product abc1 created successfully","created_at":"2000-12-31T23:59:59.000000999-02:00"}]`,
	}
}

func getTwoAuditLogs() *auditScenario {
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

	useCaseLogs := []entities.Audit{log1, log2}

	return &auditScenario{
		TestName:         "Get two products log from Audit endpoint",
		Uri:              logsV1Uri,
		MockUseCaseLogs:  useCaseLogs,
		MockUseCaseError: nil,
		ExpectStatus:     http.StatusOK,
		ExpectResponse:   `[{"request_id":"abc-123","product_id":"abc1","description":"product abc1 created successfully","created_at":"2000-12-31T23:59:59.000000999-02:00"},{"request_id":"abc-456","product_id":"abc2","description":"product abc2 created successfully","created_at":"2000-12-31T23:59:59.000000999-02:00"}]`,
	}
}

func getErrorFromUseCase() *auditScenario {
	return &auditScenario{
		TestName:         "Get error from Audit endpoint",
		Uri:              logsV1Uri,
		MockUseCaseLogs:  []entities.Audit{},
		MockUseCaseError: errorx.NewErrorx(http.StatusInternalServerError, errors.New("some-use_case-error")),
		ExpectStatus:     http.StatusInternalServerError,
		ExpectResponse:   `{"message":"some-use_case-error"}`,
	}
}

func createProductSuccessfully() *productScenario {
	product := entities.Product{
		Code:          "abc1",
		Name:          "banana",
		TotalStorage:  25,
		CorteStorage:  5,
		OriginalPrice: 10,
		SalePrice:     7,
		CreatedAt:     mockData,
		UpdatedAt:     mockData,
	}

	return &productScenario{
		TestName:          "Creates a product",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": 25,"corte": 5},"price_from": 10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockUseCase: true,
		MockUseCaseResult: product,
		MockUseCaseError:  nil,
		ExpectStatus:      http.StatusCreated,
		ExpectResponse:    `[{"code":"abc1","name":"banana","storage":{"total":25,"corte":5,"disponivel":20},"price_from":10,"price_to":7,"created_at":"2000-12-31T23:59:59.000000999-02:00","updated_at":"2000-12-31T23:59:59.000000999-02:00"}]`,
	}
}

func receivesDuplicityError() *productScenario {
	return &productScenario{
		TestName:          "Receives a duplicity code error when try to creates a new product",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": 25,"corte": 5},"price_from": 10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockUseCase: true,
		MockUseCaseResult: entities.Product{},
		MockUseCaseError:  errorx.NewErrorx(http.StatusBadRequest, errors.New("error 1062: Duplicate entry 'abc1' for key 'code'")),
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"error 1062: Duplicate entry 'abc1' for key 'code'"}`,
	}
}

func receivesRepositoryError() *productScenario {
	return &productScenario{
		TestName:          "Receives a repository error when try to creates a new product",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": 25,"corte": 5},"price_from": 10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockUseCase: true,
		MockUseCaseResult: entities.Product{},
		MockUseCaseError:  errorx.NewErrorx(http.StatusInternalServerError, errors.New("some-db-error")),
		ExpectStatus:      http.StatusInternalServerError,
		ExpectResponse:    `{"message":"some-db-error"}`,
	}
}

func receivesErrorFromBodyWithEmptyCode() *productScenario {
	return &productScenario{
		TestName:          "Receives a validation error by send a request body without code param",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "","name": "banana","storage": {"total": 25,"corte": 5},"price_from": 10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockLoggErr: true,
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"incorrect value(s) received.","fields":["code must not be empty"]}`,
	}
}

func receivesErrorFromBodyWithEmptyName() *productScenario {
	return &productScenario{
		TestName:          "Receives a validation error by send a request body without name param",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "","storage": {"total": 25,"corte": 5},"price_from": 10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockLoggErr: true,
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"incorrect value(s) received.","fields":["name must not be empty"]}`,
	}
}

func receivesErrorFromBodyWithNegativeTotal() *productScenario {
	return &productScenario{
		TestName:          "Receives a validation error by send a request body with negative storage.total param",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": -25,"corte": 5},"price_from": 10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockLoggErr: true,
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"json: cannot unmarshal number -25 into Go struct field StorageInputDto.storage.total of type uint"}`,
	}
}

func receivesErrorFromBodyWithNegativeCorte() *productScenario {
	return &productScenario{
		TestName:          "Receives a validation error by send a request body with negative storage.corte param",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": 25,"corte": -5},"price_from": 10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockLoggErr: true,
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"json: cannot unmarshal number -5 into Go struct field StorageInputDto.storage.corte of type uint"}`,
	}
}

func receivesErrorFromBodyWithZeroPriceFromAndHigherPriceTo() *productScenario {
	return &productScenario{
		TestName:          "Receives a validation error by send a request body without name param",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": 25,"corte": 5},"price_from": 0,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockLoggErr: true,
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"incorrect value(s) received.","fields":["price_from must not be lower than price_to"]}`,
	}
}

func receivesErrorFromBodyWithNegativePriceFromAndHigherPriceTo() *productScenario {
	return &productScenario{
		TestName:          "Receives a validation error by send a request body without name param",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": 25,"corte": 5},"price_from": -10,"price_to": 7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockLoggErr: true,
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"incorrect value(s) received.","fields":["price_from must not be less than zero","price_from must not be lower than price_to"]}`,
	}
}

func receivesErrorFromBodyWithNegativePriceTo() *productScenario {
	return &productScenario{
		TestName:          "Receives a validation error by send a request body without name param",
		Uri:               productsV1Uri,
		ReqBody:           `{"code": "abc1","name": "banana","storage": {"total": 25,"corte": 5},"price_from": 10,"price_to": -7}`,
		ReqMethod:         http.MethodPost,
		ShouldMockLoggErr: true,
		ExpectStatus:      http.StatusBadRequest,
		ExpectResponse:    `{"message":"incorrect value(s) received.","fields":["price_to must not be less than zero"]}`,
	}
}

func getFullUrl(uri string, pathParam string) (fullUrl string) {
	fullUrl = fmt.Sprintf("%s%s", baseUrl, uri)

	if pathParam != "" {
		fullUrl += fmt.Sprintf("/%s", pathParam)
	}

	return
}

func executeRequest(response *httptest.ResponseRecorder, method, requestUrl, body string, router *gin.Engine) {
	req, _ := http.NewRequest(method, requestUrl, bytes.NewBuffer([]byte(body)))
	router.ServeHTTP(response, req)
}
