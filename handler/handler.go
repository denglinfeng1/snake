package handler

import (
	"encoding/json"
	"net/http"
	"snake/model"
)

type server struct {

}

type route struct {

}

func (r *route)get(path, name string, handler func(http.ResponseWriter, *http.Request))  {

}

func JSON(w http.ResponseWriter, code int, resp interface{}) {
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func (s *server) HTTPRoute(r *route) {
	r.get("/game/snake/join", "join", nil)
	r.get("/game/snake/throw", "throw", nil)
	r.get("/game/snake/game_info", "game_info", nil)
	r.get("/game/snake/replay", "replay", func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query()["game_id"][0]
		game := model.GetGameById(id)
		acs := game.Actions
		JSON(writer, 200, acs)
		return
	})
}
