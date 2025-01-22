package utility

import (
	"github.com/SuperCodingTeam/model"
)

type Response struct {
	Message    string                 `json:"message"`
	StatusCode uint                   `json:"code"`
	Token      string                 `json:"token"`
	Error      *model.BookPocketError `json:"error"`
}
