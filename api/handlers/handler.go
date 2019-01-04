package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jmoiron/sqlx"
	"shortener-api/models"
	"shortener-api/restapi/operations/user"
)

type Handler struct {
	Connect *sqlx.DB
}

func (h *Handler) hitError(err error, errorCode int) middleware.Responder {
	return user.NewCreateUserBadRequest().WithPayload(
		&models.Error{
			Code:    swag.Int64(int64(errorCode)),
			Message: err.Error(),
		},
	)
}
