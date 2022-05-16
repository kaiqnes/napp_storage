package traceability

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Info(ctx *gin.Context, message string) {
	fmt.Printf("[INFO] [%v] [RequestID:%s] %s\n", time.Now().Format(time.RFC3339), GetRequestID(ctx), message)
}

func Error(ctx *gin.Context, message string) {
	fmt.Printf("[ERROR] [%v] [RequestID:%s] %s\n", time.Now().Format(time.RFC3339), GetRequestID(ctx), message)
}
