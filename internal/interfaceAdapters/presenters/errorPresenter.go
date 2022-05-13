package presenters

import (
	"github.com/gin-gonic/gin"
	"storage/internal/interfaceAdapters/dto"
)

type ErrorPresenter interface {
	PresentError(ctx *gin.Context, err error, statusCode int)
}

type errorPresenter struct{}

func NewErrorPresenter() ErrorPresenter {
	return &errorPresenter{}
}

func (pbl *errorPresenter) PresentError(ctx *gin.Context, err error, statusCode int) {
	errResponse := dto.ErrorOutputDto{Message: err.Error()}
	ctx.JSON(statusCode, errResponse)
}
