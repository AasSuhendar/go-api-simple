package routers

import (
	"go-simple-api/models"
)

type ResponseGetUser struct {
	Status  bool          `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}
