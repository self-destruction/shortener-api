package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ivahaev/go-logger"
	"shortener-api/models"
	"shortener-api/restapi/operations/statistic"
)

func (h *Handler) GetCurrentUser(params statistic.GetCurrentUserParams, principal interface{}) middleware.Responder {
	// transform interface{} to User{}
	user := principal.(*models.User)
	logger.JSON(user)
	return statistic.NewGetCurrentUserOK().WithPayload(user)
}
