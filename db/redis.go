package db

import (
	"encoding/json"
	"github.com/Avlek/voiceChessApi/models"
)
import "github.com/go-redis/redis"

var client *redis.Client

var preparedMoves = []models.MoveObject{
	models.MoveObject{Player: "w", Move: "e2e4"},
	models.MoveObject{Player: "b", Move: "e7e5"},
	models.MoveObject{Player: "w", Move: "g1f3"},
	models.MoveObject{Player: "b", Move: "b8c6"},
	models.MoveObject{Player: "w", Move: "f1c4"},
	models.MoveObject{Player: "b", Move: "d8e7"},
}

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	InitPreparedMoves("test")
}

func InitPreparedMoves(game string) {

	for _, value := range preparedMoves {
		SaveMove(game, value)
	}
}

func getMoveKey(game string) string {
	return "game_" + game
}

func SaveMove(game string, m models.MoveObject) error {
	b, _ := json.Marshal(m)

	err := client.RPush(getMoveKey(game), b).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetMoves(game string) (moves models.MoveObjects, err error) {
	rows, err := client.LRange(getMoveKey(game), 0, -1).Result()
	if err != nil {
		return moves, err
	}

	ms := []models.MoveObject{}
	for _, val := range rows {
		var m = models.MoveObject{}
		err = json.Unmarshal([]byte(val), &m)
		ms = append(ms, m)
	}

	return ms, err
}
