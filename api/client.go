package api

import (
	"errors"
	"github.com/Avlek/voiceChessApi/db"
	"github.com/Avlek/voiceChessApi/models"
	"regexp"
	"strings"
)

func Move(m models.MoveObject) (string, error) {

	m.Move = strings.Replace(m.Move, " ", "", -1)

	if m.Move == "" {
		return "", errors.New("Пустые данные")
	}

	if b, err := regexp.MatchString("[a-h][1-8][a-h][1-8]", m.Move); !b || err != nil {
		return "", errors.New("Некорректные данные")
	}

	status, err := validateMove(m)

	if err != nil {
		return "", err
	}
	db.SaveMove("test", m)

	return status, nil
}
