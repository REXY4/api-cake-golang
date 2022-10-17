package helper

import "strings"

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors"`
	Data       interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(statusCode int, message string, data interface{}) Response {
	res := Response{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	return res
}

func BuildErrorResponse(statusCode int, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:     "Failed",
		StatusCode: statusCode,
		Errors:     splittedError,
		Message:    message,
		Data:       data,
	}
	return res
}
