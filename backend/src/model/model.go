package model

import (
	"fmt"
	"time"
)

// State object definitions:
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
	Name   string
	Node   NodeID
	ID     PlayerID
	Role   Role
	Color  string // TODO: figure out best way to encode this
	admin  bool
	checks []Check
}

type Check struct {
	name         string
	guessedValue string
	correct      bool
}

type Graph struct {
	nodes  map[NodeID]Node
	groups map[GroupID]Group
}

type Node struct {
	ID     NodeID
	Player PlayerID
	Name   string
	Color  string
	Groups []GroupID
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
	Value    string
	Possible []string
}

type Facts struct {
	Real      map[string]Fact
	Submitted map[string]string
}

// State update structs and functions
type EventName struct {
	Event string
}

type UpdatePlayer struct {
	EventName
	Player Player
}

func (s *State) UpdatePlayer(newPlayer Player) {
	oldPlayer := s.Players[newPlayer.ID]
	*oldPlayer = newPlayer
}

type RemovePlayer struct {
	EventName
	PlayerID PlayerID
}

func (s *State) RemovePlayer(id PlayerID) error {
	delete(s.Players, id)
	return nil
}

// TODO: Not sure what this should do
func (s *State) StopGame() error {
	return nil
}

type CheckFact struct {
	EventName
	PlayerID PlayerID
	Field    string
	Value    string
}

func (s *State) CheckFact(check *CheckFact) error {
	guessingPlayer, ok := s.Players[check.PlayerID]
	if !ok {
		return fmt.Errorf("no existing players with PlayerID %v", check.PlayerID)
	}
	fact, ok := s.Facts.Real[check.Field]
	if !ok {
		return fmt.Errorf("no fact with name %v", check.Field)
	}
	guessingPlayer.checks = append(guessingPlayer.checks, Check{name: check.Field, guessedValue: check.Value, correct: check.Value == fact.Value})
	return nil
}

func (s *State) StartPregame() {
	s.Game.state = LOBBY
	// TODO Do we need or want this to do other state cleanup?
}

type StartGame struct {
	EventName
	StopTime time.Time
}

func (s *State) StartGame(stopTime time.Time) {
	s.Game.state = PLAYING
	s.Game.startTime = time.Now()
	s.Game.stopTime = stopTime
}

type SubmitFacts struct {
	EventName
	Submission map[string]string
}

func (s *State) SubmitFacts(submission map[string]string) {
	s.Facts.Submitted = submission
}

type SendMessage struct {
	EventName
	GroupID GroupID
	Sender  NodeID
	Text    string
}

func (s *State) SendMessage(newMsg *SendMessage) error {
	group, ok := s.Graph.groups[newMsg.GroupID]
	if !ok {
		return fmt.Errorf("group with id %v does not exist", newMsg.GroupID)
	}
	group.Messages = append(group.Messages, Message{Sender: newMsg.Sender, Timestamp: time.Now(), Text: newMsg.Text})
	return nil
}
