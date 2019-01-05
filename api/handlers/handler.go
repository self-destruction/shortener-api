package handlers

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/ivahaev/go-logger"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"shortener-api/models"
)

type Handler struct {
	Connect *sqlx.DB
}

func (h *Handler) Authorize(user string, pass string) (interface{}, error) {
	userDB := &models.User{}
	query := "SELECT id, username, hash, email, createdAt AS dateCreated FROM `shortener`.`user`WHERE username = ? AND hash = ? LIMIT 1"
	err := h.Connect.Get(userDB, query, user, pass)
	if err != nil {
		logger.Notice(err)
		return nil, errors.Unauthenticated(fmt.Sprintf("this username and password"))
	}

	return userDB, nil
}

func (h *Handler) getHash(password string) string {
	// Hash & salt the user password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logger.Debug(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
