package traceability

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context, message string) {
	fmt.Println("[INFO]", GetRequestID(ctx), message)
}

//func Infof(ctx *gin.Context, template string, args ...interface{}) {
//	msgWithReqID := fmt.Sprintf("%s %s", GetRequestID(ctx), template)
//	fmt.Printf(fmt.Sprintf(msgWithReqID, args))
//}
//
//func Info(ctx *gin.Context, message string) {
//	fmt.Println(fmt.Sprintf("%s %s", GetRequestID(ctx), message))
//}
