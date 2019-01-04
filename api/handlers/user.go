package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ivahaev/go-logger"
	"shortener-api/models"
	"shortener-api/restapi/operations/user"
)

func (h *Handler) CreateUser(params user.CreateUserParams) middleware.Responder {
	userRequest := params.Body

	stmt, err := h.Connect.Prepare("INSERT `shortener`.`user` SET username=?, hash=?, email=?, timezone=?, language=?")
	if err != nil {
		logger.Debug(err)
		return h.hitError(err, user.CreateUserBadRequestCode)
	}

	// insert user into db
	res, err := stmt.Exec(userRequest.Username, userRequest.Password, userRequest.Email, userRequest.Timezone, userRequest.Language)
	if err != nil {
		logger.Debug(err)
		return h.hitError(err, user.CreateUserBadRequestCode)
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		logger.Debug(err)
		return h.hitError(err, user.CreateUserBadRequestCode)
	}

	// select user by last insert id for response
	userDB := &models.User{}
	query := "SELECT id, username, hash, email, timezone, language, createdAt AS dateCreated FROM `shortener`.`user` WHERE id=? LIMIT 1"
	err = h.Connect.Get(userDB, query, insertId)
	if err != nil {
		logger.Debug(err)
		return h.hitError(err, user.CreateUserBadRequestCode)
	}

	logger.JSON(userDB)
	return user.NewCreateUserOK().WithPayload(userDB)
}
