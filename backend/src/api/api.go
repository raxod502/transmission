package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raxod502/transmission/backend/src/model"
)

func launchMissiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Missiles launched\n"))
}

func Setup(r *mux.Router) {
	r.HandleFunc("/launch", launchMissiles)
}

func handleEvent(payload []byte, state model.State) error {
	event := &model.EventName{}
	json.Unmarshal(payload, event)
	eventHandlers := map[string]func([]byte) error{
		"updatePlayer": func(payload []byte) error {
			message := &model.UpdatePlayer{}
			json.Unmarshal(payload, message)
			return state.UpdatePlayer(message.Player)
		},
		"removePlayer": func(payload []byte) error {
			message := &model.RemovePlayer{}
			json.Unmarshal(payload, message)
			state.RemovePlayer(message.PlayerID)
			return nil
		},
		"stopGame": func(payload []byte) error {
			return state.StopGame()
		},
		"checkFact": func(payload []byte) error {
			message := &model.CheckFact{}
			json.Unmarshal(payload, message)
			return state.CheckFact(message)
		},
		"startPreGame": func(payload []byte) error {
			state.StartPregame()
			return nil
		},
		"startGame": func(payload []byte) error {
			message := &model.StartGame{}
			json.Unmarshal(payload, message)
			state.StartGame(message.StopTime)
			return nil
		},
		"submitFacts": func(payload []byte) error {
			message := &model.SubmitFacts{}
			json.Unmarshal(payload, message)
			state.SubmitFacts(message.Submission)
			return nil
		},
		"sendMessage": func(payload []byte) error {
			message := &model.SendMessage{}
			json.Unmarshal(payload, message)
			state.SendMessage(message)
			return nil
		},
	}
	handler, ok := eventHandlers[event.Event]
	if !ok {
		return fmt.Errorf("event type %v does not have a handler", event.Event)
	}
	err := handler(payload)
	return err
}
