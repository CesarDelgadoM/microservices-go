package common

type Response interface {
}

type responseData struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"error_code"`
}

func ResponseData(data interface{}, msg string, statusCode int) Response {
	return responseData{
		Data:       data,
		Message:    msg,
		StatusCode: statusCode,
	}
}

type responseError struct {
	ErrorMessage string `json:"error_message"`
	StatusCode   int    `json:"error_code"`
}

func ResponseError(errorMessage string, statusCode int) Response {
	return responseError{
		ErrorMessage: errorMessage,
		StatusCode:   statusCode,
	}
}
