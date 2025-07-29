package handler

import (
	"github.com/doruo/falloutdle/internal/character"
)

type CharacterHandler struct {
	characterService *character.Service
}

func NewCharacterHandler(ch *character.Service) *CharacterHandler {
	return &CharacterHandler{
		characterService: ch,
	}
}
