package model

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type PlayerID string
type NodeID string
type GroupID string

type Role string

const (
	HQ  Role = "HQ"
	TD  Role = "TD"
	BAD Role = "Baddie"
	TE  Role = "Train Expert"
	FC  Role = "Fact Checker"
)

type Gamestate string

const (
	LOBBY      Gamestate = "Lobby"
	PLAYING    Gamestate = "Playing"
	SUBMISSION Gamestate = "Submission"
	RESULTS    Gamestate = "Results"
)

type State struct {
	Game    Game
	Players map[PlayerID]*Player
	Graph   Graph
	Facts   Facts
}

type Game struct {
	state     Gamestate
	startTime time.Time
	stopTime  time.Time
}

type Player struct {
	Name  string
	Node  NodeID
	Role  Role
	Color string // TODO: figure out best way to encode this
	admin bool
}

type Graph struct {
	nodes  map[NodeID][]GroupID
	groups map[GroupID]Group
}

type Group struct {
	Messages []Message
}

type Message struct {
	Sender    NodeID
	Text      string
	Timestamp time.Time
}

type Fact struct {
	Value          string
	PossibleValues []string
	Checked        map[string]bool
}

type Facts map[string]Fact

func validateEventName(data map[string]string, expectedName string) error {
	eventName, ok := data["event"]
	if !ok || eventName != expectedName {
		return fmt.Errorf(
			"called method on payload without an appropriate event field. Expected %v, but got %v",
			expectedName,
			eventName,
		)
	}
	return nil
}

func (s State) updatePlayer(data map[string]string) error {
	err := validateEventName(data, "updatePlayer")
	if err != nil {
		return err
	}

	playerID, ok := data["playerID"]
	if !ok {
		return errors.New("missing playerID field")
	}

	admin, err := strconv.ParseBool(data["playerAdmin"])
	if err != nil {
		return fmt.Errorf("Failed to parse playerAdmin value into bool with error %v", err)
	}
	s.Players[PlayerID(playerID)] = &Player{
		Name:  data["playerName"],
		Node:  NodeID(data["playerNode"]),
		Role:  Role(data["playerRole"]),
		Color: data["playerColor"],
		admin: admin,
	}
	return nil
}

func (s State) removePlayer(data map[string]string) error {
	err := validateEventName(data, "removePlayer")
	if err != nil {
		return err
	}

	playerID, ok := data["playerID"]
	if !ok {
		return errors.New("missing playerID field")
	}
	delete(s.Players, PlayerID(playerID))
	return nil
}
