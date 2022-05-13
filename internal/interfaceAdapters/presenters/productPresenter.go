package presenters

import (
	"github.com/gin-gonic/gin"
)

type ProductPresenter interface {
	PresentProduct(ctx *gin.Context, err error, statusCode int)
}

type productPresenter struct {
	errorPresenter ErrorPresenter
}

func NewProductPresenter(errorPresenter ErrorPresenter) ProductPresenter {
	return &productPresenter{errorPresenter: errorPresenter}
}

func (pbl *productPresenter) PresentProduct(ctx *gin.Context, err error, statusCode int) {

}
