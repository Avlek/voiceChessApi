package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Avlek/voiceChessApi/api"
	"github.com/Avlek/voiceChessApi/db"
	"github.com/Avlek/voiceChessApi/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var moves = []string{}

type ErrStruct struct {
	Error string `json:"error"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	p := models.Color(r.Header.Get("user")[0])

	if p == "" {
		fmt.Fprint(w, errors.New("Игрок не найден"))
		return
	}

	if ps.ByName("method") == "move" {
		decoderBody := json.NewDecoder(r.Body)
		var m models.MoveObject
		err := decoderBody.Decode(&m)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		m.Player = p
		err = api.Move(m)
		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			fmt.Fprint(w, fmt.Sprintf("Ваш ход %s", m.Move))
		}
	}
}

func Web(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if p.ByName("method") == "getmoves" {
		moves, err := api.GetMovesResponse("test")
		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			fmt.Fprint(w, moves)
		}
	}
}

func main() {

	db.InitRedis()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/web/:method", Web)
	router.POST("/api/:method", Api)

	log.Fatal(http.ListenAndServe(":8089", router))
}
