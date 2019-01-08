package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ivahaev/go-logger"
	"golang.org/x/crypto/bcrypt"
	"shortener-api/models"
	"shortener-api/restapi/operations/user"
)

func (h *Handler) LoginUser(params user.LoginUserParams) middleware.Responder {
	userRequest := params.Body

	userDB := &models.User{}
	query := "SELECT id, username, hash, email, createdAt AS dateCreated FROM `shortener`.`user`WHERE email = ? LIMIT 1"
	err := h.Connect.Get(userDB, query, userRequest.Email)
	if err != nil {
		logger.Debug(err)
		return hitLoginUserError(err, user.LoginUserBadRequestCode)
	}

	// compare password and hash
	err = bcrypt.CompareHashAndPassword([]byte(*userDB.Hash), []byte(*userRequest.Password))
	if err != nil {
		logger.Debug(err)
		return hitLoginUserError(err, user.LoginUserBadRequestCode)
	}

	logger.JSON(userDB)
	return user.NewLoginUserOK().WithPayload(userDB)
}

func (h *Handler) CreateUser(params user.CreateUserParams) middleware.Responder {
	userRequest := params.Body

	hash := h.getHash(*userRequest.Password)

	queryString := "INSERT `shortener`.`user` SET username=?, hash=?, email=?"
	stmt, err := h.Connect.Prepare(queryString)
	if err != nil {
		logger.Debug(err)
		return hitCreateUserError(err, user.CreateUserBadRequestCode)
	}

	// insert user into db
	res, err := stmt.Exec(userRequest.Username, hash, userRequest.Email)
	if err != nil {
		logger.Debug(err)
		return hitCreateUserError(err, user.CreateUserBadRequestCode)
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		logger.Debug(err)
		return hitCreateUserError(err, user.CreateUserBadRequestCode)
	}

	// select user by last insert id for response
	userDB := &models.User{}
	query := "SELECT id, username, hash, email, createdAt AS dateCreated FROM `shortener`.`user` WHERE id=? LIMIT 1"
	err = h.Connect.Get(userDB, query, insertId)
	if err != nil {
		logger.Debug(err)
		return hitCreateUserError(err, user.CreateUserBadRequestCode)
	}

	logger.JSON(userDB)
	return user.NewCreateUserOK().WithPayload(userDB)
}

func hitCreateUserError(err error, errorCode int) middleware.Responder {
	return user.NewCreateUserBadRequest().WithPayload(
		&models.Error{
			Code:    swag.Int64(int64(errorCode)),
			Message: err.Error(),
		},
	)
}

func hitLoginUserError(err error, errorCode int) middleware.Responder {
	return user.NewCreateUserBadRequest().WithPayload(
		&models.Error{
			Code:    swag.Int64(int64(errorCode)),
			Message: err.Error(),
		},
	)
}
