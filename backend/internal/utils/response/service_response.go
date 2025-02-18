package response

type Response struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func Success(message string, statusCode int, data interface{}) Response {
	return Response{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	}
}

func Error(message string, statusCode int) Response {
	return Response{
		Message:    message,
		StatusCode: statusCode,
	}
}
