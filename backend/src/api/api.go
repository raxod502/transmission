package api

import (
	"encoding/json"
	"fmt"

	"github.com/raxod502/transmission/backend/src/model"
)

func HandleEvent(payload []byte, state *model.State) error {
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
			return state.StartGame(message.StopTime)
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
		"updateNode": func(payload []byte) error {
			message := &model.UpdateNode{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.UpdateNode(message)
			return nil
		},
		"removeNode": func(payload []byte) error {
			message := &model.RemoveNode{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.RemoveNode(message.NodeID)
			return nil
		},
		"updateGroup": func(payload []byte) error {
			message := &model.UpdateGroup{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.UpdateGroup(message)
			return nil
		},
		"removeGroup": func(payload []byte) error {
			message := &model.RemoveGroup{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.RemoveGroup(message.GroupID)
			return nil
		},
		"updateRealFactPossibilities": func(payload []byte) error {
			message := &model.UpdateRealFactPossibilities{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.UpdateRealFactPossibilities(message)
			return nil
		},
		"setRealFacts": func(payload []byte) error {
			message := &model.SetRealFacts{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			state.SetRealFacts(message)
			return nil
		},
		"addKnownRole": func(payload []byte) error {
			message := &model.AddKnownRole{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			return state.AddKnownRole(message)
		},
		"addKnownFact": func(payload []byte) error {
			message := &model.AddKnownFact{}
			err := json.Unmarshal(payload, message)
			if err != nil {
				return err
			}
			return state.AddKnownFact(message)
		},
	}
	handler, ok := eventHandlers[event.Event]
	if !ok {
		return fmt.Errorf("event type %v does not have a handler", event.Event)
	}
	err = handler(payload)
	if err != nil {
		return err
	}

	return err
}
