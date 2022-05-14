package helper

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
