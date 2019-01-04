package handlers

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/ivahaev/go-logger"
	"github.com/jmoiron/sqlx"
	"shortener-api/models"
)

type Handler struct {
	Connect *sqlx.DB
}

func (h *Handler) Authorize(user string, pass string) (interface{}, error) {
	userDB := &models.User{}
	query := "SELECT id, username, hash, email, timezone, language, createdAt AS dateCreated FROM `shortener`.`user`WHERE username = ? AND hash = ? LIMIT 1"
	err := h.Connect.Get(userDB, query, user, pass)
	if err != nil {
		logger.Notice(err)
		return nil, errors.Unauthenticated(fmt.Sprintf("username - %s", user))
	}

	return userDB, nil
}
