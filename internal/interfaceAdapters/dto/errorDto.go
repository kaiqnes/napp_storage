package dto

type ErrorOutputDto struct {
	Message string `json:"message"`
}

type ErrorsOutputDto struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields"`
}
