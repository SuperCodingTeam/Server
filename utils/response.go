package utils

import (
	"github.com/SuperCodingTeam/models"
)

type Response struct {
	Message    string                  `json:"message"`
	StatusCode uint                    `json:"code"`
	Token      string                  `json:"token"`
	Error      *models.BookPocketError `json:"error"`
}
