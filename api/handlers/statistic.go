package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ivahaev/go-logger"
	"shortener-api/models"
	"shortener-api/restapi/operations/statistic"
)

func (h *Handler) GetLinkInfo(params statistic.GetLinkInfoParams, principal interface{}) middleware.Responder {
	return middleware.NotImplemented("operation statistic.GetLinkInfo has not yet been implemented")
	//
	//	// select link by hash for response
	//	linkDB := &models.Link{}
	//	query := "SELECT id, userId, shortUrl, fullUrl, createdAt AS dateCreated FROM `shortener`.`link` WHERE shortUrl=? LIMIT 1"
	//	err := h.Connect.Get(linkDB, query, params.Hash)
	//	if err != nil {
	//		logger.Debug(err)
	//		return link.NewGetLinkNotFound().WithPayload(
	//			&models.Error{
	//				Code:    swag.Int64(int64(link.GetLinkNotFoundCode)),
	//				Message: err.Error(),
	//			},
	//		)
	//	}
	//
	//	logger.JSON(linkDB)
	//	return statistic.NewGetLinkInfoOK().WithPayload(linkDB)
}

func (h *Handler) GetCurrentUser(params statistic.GetCurrentUserParams, principal interface{}) middleware.Responder {
	// transform interface{} to User{}
	user := principal.(*models.User)
	logger.JSON(user)
	return statistic.NewGetCurrentUserOK().WithPayload(user)
}
