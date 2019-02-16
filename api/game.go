package api

import (
	"errors"
	"github.com/Avlek/voiceChessApi/models"
	"github.com/andrewbackes/chess/game"
	"github.com/andrewbackes/chess/position/move"
)

var g *game.Game

func StartGame() {

	g = game.New()

}

func validateMove(m models.MoveObject) (string, error) {

	if m.Player != uint8(g.ActiveColor()) {
		return "", errors.New("Сейчас не ваш ход")
	}

	s := move.Parse(m.Move)

	status, err := g.MakeMove(s)

	if err != nil {
		return "", errors.New("Невозможный ход")
	}

	return status.String(), nil
}
