package helper

import "github.com/go-playground/validator/v10"

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"` //field data biar fleksibel
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{}
	meta.Message = message
	meta.Code = code
	meta.Status = status

	jsonResponse := Response{}
	jsonResponse.Meta = meta
	jsonResponse.Data = data

	return jsonResponse
}

func ErrorValidationResponse(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
