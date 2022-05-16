package errorx

type errorx struct {
	statusCode int
	error      error
}

type Errorx interface {
	GetStatusCode() int
	GetError() error
}

func NewErrorx(statusCode int, err error) Errorx {
	return &errorx{
		statusCode: statusCode,
		error:      err,
	}
}

func (e errorx) GetStatusCode() int {
	return e.statusCode
}

func (e errorx) GetError() error {
	return e.error
}
