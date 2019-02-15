package api

import (
	"errors"
	"github.com/Avlek/voiceChessApi/db"
	"github.com/Avlek/voiceChessApi/models"
)

func Move(m models.MoveObject) error {
	if m.Move == "" {
		return errors.New("Пустые данные")
	}

	db.SaveMove("test", m)

	return nil
}
