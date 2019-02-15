package api

import (
	"errors"
	"github.com/Avlek/voiceChessApi/models"
)

func Move(m models.MoveObject) error {
	if m.Move == "" {
		return errors.New("Пустые данные")
	}

	return nil
}
