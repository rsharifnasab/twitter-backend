package handler

import (
	"github.com/rsharifnasab/twitter-backend/article"
	"github.com/rsharifnasab/twitter-backend/user"
)

type Handler struct {
	userStore    user.Store
	articleStore article.Store
}

func NewHandler(us user.Store, as article.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
	}
}
