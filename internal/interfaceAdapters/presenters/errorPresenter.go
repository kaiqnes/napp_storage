package presenters

import (
	"github.com/gin-gonic/gin"
	"storage/internal/interfaceAdapters/dto"
)

type ErrorPresenter interface {
	PresentError(ctx *gin.Context, err error, statusCode int)
	PresentErrors(ctx *gin.Context, errDto dto.ErrorsOutputDto, statusCode int)
}

type errorPresenter struct{}

func NewErrorPresenter() ErrorPresenter {
	return &errorPresenter{}
}

func (presenter *errorPresenter) PresentError(ctx *gin.Context, err error, statusCode int) {
	errResponse := dto.ErrorOutputDto{Message: err.Error()}
	ctx.JSON(statusCode, errResponse)
}

func (presenter *errorPresenter) PresentErrors(ctx *gin.Context, errDto dto.ErrorsOutputDto, statusCode int) {
	ctx.JSON(statusCode, errDto)
}
