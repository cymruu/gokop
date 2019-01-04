package models

import "fmt"

type ErrorResponse struct {
	ErrorObject struct {
		Message string `json:"message"`
		Code    uint16 `json:"code"`
	} `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s [%d]", e.ErrorObject.Message, e.ErrorObject.Code)
}
