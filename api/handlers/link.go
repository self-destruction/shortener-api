package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/ivahaev/go-logger"
	"math/rand"
	"shortener-api/models"
	"shortener-api/restapi/operations/link"
	"time"
)

func (h *Handler) GetLink(params link.GetLinkParams) middleware.Responder {
	// select link by hash for response
	linkDB := &models.Link{}
	query := "SELECT id, userId, shortUrl, fullUrl, createdAt AS dateCreated FROM `shortener`.`link` WHERE shortUrl=? LIMIT 1"
	err := h.Connect.Get(linkDB, query, params.Hash)
	if err != nil {
		logger.Debug(err)
		return link.NewGetLinkNotFound().WithPayload(
			&models.Error{
				Code:    swag.Int64(int64(link.GetLinkNotFoundCode)),
				Message: err.Error(),
			},
		)
	}

	logger.JSON(linkDB)
	return link.NewGetLinkMovedPermanently().WithLocation(*linkDB.FullURL)
}

func (h *Handler) GetLinks(params link.GetLinksParams, principal interface{}) middleware.Responder {
	user := principal.(*models.User)

	// select array links by user id for response
	query := "SELECT id, userId, shortUrl, fullUrl, createdAt AS dateCreated FROM `shortener`.`link` WHERE userId=?"
	rows, err := h.Connect.Queryx(query, user.ID)
	if err != nil {
		logger.Debug(err)
		return link.NewCreateShortLinkBadRequest().WithPayload(
			&models.Error{
				Code:    swag.Int64(int64(link.GetLinkNotFoundCode)),
				Message: err.Error(),
			},
		)
	}

	linkBody := link.GetLinksOKBody{}

	for rows.Next() {
		//TODO maybe may reuse this structure?
		linkDB := &models.Link{}
		err = rows.StructScan(linkDB)
		if err != nil {
			logger.Debug(err)
			return link.NewCreateShortLinkBadRequest().WithPayload(
				&models.Error{
					Code:    swag.Int64(int64(link.GetLinkNotFoundCode)),
					Message: err.Error(),
				},
			)
		}
		linkBody.Links = append(linkBody.Links, linkDB)
	}

	logger.Debug(linkBody.Links)
	return link.NewGetLinksOK().WithPayload(&linkBody)
}

func (h *Handler) CreateShortLink(params link.CreateShortLinkParams, principal interface{}) middleware.Responder {
	userRequest := params.Body
	user := principal.(*models.User)

	queryString := "INSERT `shortener`.`link` SET shortUrl=?, fullUrl=?, userId=?"
	stmt, err := h.Connect.Prepare(queryString)
	if err != nil {
		logger.Debug(err)
		return hitCreateShortUrlError(err, link.CreateShortLinkBadRequestCode)
	}

	shortUrl := RandStringBytesMaskImprSrc(8)

	// insert user into db
	res, err := stmt.Exec(shortUrl, userRequest.URL, user.ID)
	if err != nil {
		logger.Debug(err)
		return hitCreateShortUrlError(err, link.CreateShortLinkBadRequestCode)
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		logger.Debug(err)
		return hitCreateShortUrlError(err, link.CreateShortLinkBadRequestCode)
	}

	// select link by last insert id for response
	linkDB := &models.Link{}
	query := "SELECT id, userId, shortUrl, fullUrl, createdAt AS dateCreated FROM `shortener`.`link` WHERE id=? LIMIT 1"
	err = h.Connect.Get(linkDB, query, insertId)
	if err != nil {
		logger.Debug(err)
		return hitCreateShortUrlError(err, link.CreateShortLinkBadRequestCode)
	}

	logger.JSON(linkDB)
	return link.NewCreateShortLinkOK().WithPayload(linkDB)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func hitCreateShortUrlError(err error, errorCode int) middleware.Responder {
	return link.NewCreateShortLinkBadRequest().WithPayload(
		&models.Error{
			Code:    swag.Int64(int64(errorCode)),
			Message: err.Error(),
		},
	)
}
