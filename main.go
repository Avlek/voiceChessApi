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

type ErrStruct struct {
	Error string `json:"error"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	p := r.Header.Get("user")
	var ap uint8

	if p == "white" {
		ap = 0
	} else if p == "black" {
		ap = 1
	} else {
		fmt.Fprint(w, errors.New("Игрок не найден"))
		return
	}

	if ps.ByName("method") == "move" {
		decoderBody := json.NewDecoder(r.Body)
		var m models.ClientMoveObject
		err := decoderBody.Decode(&m)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		status, err := api.Move(ap, m)
		if err != nil {
			fmt.Fprint(w, err.Error())
		} else {
			fmt.Fprint(w, fmt.Sprintf(status))
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
	} else if p.ByName("method") == "flushall" {
		db.DeleteTestMoves()
	}
}

func main() {

	db.InitRedis()

	api.StartGame()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/web/:method", Web)
	router.POST("/api/:method", Api)

	log.Fatal(http.ListenAndServe(":8089", router))
}
