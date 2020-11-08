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

func HandleEvent(payload []byte, state model.State) error {
	event := &model.EventName{}
	err := json.Unmarshal(payload, event)
	if err != nil {
		return err
	}
	eventHandlers := map[string]func([]byte) error{
		"updatePlayer": func(payload []byte) error {
			message := &model.UpdatePlayer{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.UpdatePlayer(message.Player)
			return nil
		},
		"removePlayer": func(payload []byte) error {
			message := &model.RemovePlayer{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.RemovePlayer(message.PlayerID)
			return nil
		},
		"stopGame": func(payload []byte) error {
			return state.StopGame()
		},
		"checkFact": func(payload []byte) error {
			message := &model.CheckFact{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			return state.CheckFact(message)
		},
		"startPregame": func(payload []byte) error {
			state.StartPregame()
			return nil
		},
		"startGame": func(payload []byte) error {
			message := &model.StartGame{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.StartGame(message.StopTime)
			return nil
		},
		"submitFacts": func(payload []byte) error {
			message := &model.SubmitFacts{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.SubmitFacts(message.Submission)
			return nil
		},
		"sendMessage": func(payload []byte) error {
			message := &model.SendMessage{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.SendMessage(message)
			return nil
		},
	}
	handler, ok := eventHandlers[event.Event]
	if !ok {
		return fmt.Errorf("event type %v does not have a handler", event.Event)
	}
	err = handler(payload)
	return err
}
