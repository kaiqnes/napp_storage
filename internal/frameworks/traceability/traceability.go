package traceability

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewRequestID() string {
	return fmt.Sprintf("[RequestID:%s]", uuid.New().String())
}

func GetRequestID(ctx *gin.Context) string {
	return ctx.Request.Header.Get("request-id")
}

func SetRequestID(ctx *gin.Context, requestID string) {
	ctx.Request.Header.Set("request-id", requestID)
}

func ValidateRequestID(ctx *gin.Context) {
	requestID := GetRequestID(ctx)
	if requestID == "" {
		requestID = NewRequestID()
		SetRequestID(ctx, requestID)
	}
}
