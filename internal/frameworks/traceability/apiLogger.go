package traceability

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type apiLogger struct {
}

//go:generate mockgen -source=./apiLogger.go -destination=./mocks/apiLogger_mock.go
type ApiLogger interface {
	Info(ctx *gin.Context, message string)
	Error(ctx *gin.Context, message string)
}

func NewApiLogger() ApiLogger {
	return &apiLogger{}
}

func (logg *apiLogger) Info(ctx *gin.Context, message string) {
	fmt.Printf("[INFO] [%v] [RequestID:%s] %s\n", time.Now().Format(time.RFC3339), GetRequestID(ctx), message)
}

func (logg *apiLogger) Error(ctx *gin.Context, message string) {
	fmt.Printf("[ERROR] [%v] [RequestID:%s] %s\n", time.Now().Format(time.RFC3339), GetRequestID(ctx), message)
}
