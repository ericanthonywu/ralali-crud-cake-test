package Model

type ResponseDto struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func SuccessResponse(message string, data interface{}) *ResponseDto {
	return &ResponseDto{Message: message, Data: data}
}

func ErrorResponse(message string, error interface{}) *ResponseDto {
	return &ResponseDto{Message: message, Error: error}
}

type ValidationError struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}
