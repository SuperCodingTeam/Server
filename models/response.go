package models

import "fmt"

type BookPocketError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *BookPocketError) Error() string {
	return fmt.Sprintf("%s, %s", e.Code, e.Message)
}
