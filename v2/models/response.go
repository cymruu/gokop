package models

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Data       json.RawMessage `json:"data"`
	Pagination *Pagination     `json:"pagination"`
	Error      *Error          `json:"error"`
}

type Error struct {
	Code      int    `json:"code"`
	Field     string `json:"field"`
	MessagePl string `json:"message_pl"`
	MessageEn string `json:"message_en"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Wykop API Error code: %d: %s", e.Code, e.MessageEn)
}

type Pagination struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
