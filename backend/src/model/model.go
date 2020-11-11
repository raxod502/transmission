package model

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// State object definitions:
type PlayerID string
type NodeID string
type GroupID string

type Role string

const (
	HEADQUARTERS Role = "Headquarters"
	TRAINDEPOT   Role = "Train Depot"
	DOUBLEAGENT  Role = "Double Agent"
	TRAINEXPERT  Role = "Train Expert"
	FACTCHECKER  Role = "Fact Checker"
)

type Gamestate string

const (
	LOBBY      Gamestate = "lobby"
	PLAYING    Gamestate = "playing"
	SUBMISSION Gamestate = "submission"
	RESULTS    Gamestate = "results"
)

func NewState() *State {
	return &State{
		Game: Game{
			State: PLAYING, // change this later to LOBBY
		},
		Players: map[PlayerID]*Player{},
		Graph: Graph{
			Nodes:  map[NodeID]*Node{},
			Groups: map[GroupID]*Group{},
		},
		Facts: Facts{
			Real:      map[string]*Fact{},
			Submitted: map[string]string{},
		},
	}
}

type State struct {
	Game    Game                 `json:"game"`
	Players map[PlayerID]*Player `json:"players"`
	Graph   Graph                `json:"graph"`
	Facts   Facts                `json:"facts"`
}

type Game struct {
	State     Gamestate  `json:"state"`
	StartTime *time.Time `json:"startTime"`
	StopTime  *time.Time `json:"stopTime"`
}

type Player struct {
	Name string   `json:"name"`
	Node NodeID   `json:"node"`
	ID   PlayerID `json:"id"`
	Role Role     `json:"role"`
	// TODO: figure out best way to encode colors
	Color  string  `json:"color"`
	checks []Check `json:"checks"`
}

type Check struct {
	Name         string `json:"name"`
	GuessedValue string `json:"guessedValue"`
	Correct      bool   `json:"correct"`
}

type Graph struct {
	Nodes  map[NodeID]*Node   `json:"nodes"`
	Groups map[GroupID]*Group `json:"groups"`
}

type Node struct {
	ID     NodeID    `json:"id"`
	Player PlayerID  `json:"player"`
	Name   string    `json:"name"`
	Color  string    `json:"color"`
	Groups []GroupID `json:"groups"`
}

type Group struct {
	ID       GroupID   `json:"id"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Sender    NodeID    `json:"sender"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type Fact struct {
	Value    string   `json:"value"`
	Possible []string `json:"possible"`
}

type Facts struct {
	Real      map[string]*Fact  `json:"real"`
	Submitted map[string]string `json:"submitted"`
}

// State update structs and functions
type EventName struct {
	Event string `json:"event"`
}

type UpdatePlayer struct {
	EventName
	Player Player `json:"player"`
}

func (s *State) UpdatePlayer(newPlayer Player) {
	s.Players[newPlayer.ID] = &newPlayer
}

type RemovePlayer struct {
	EventName
	PlayerID PlayerID `json:"playerID"`
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
	PlayerID PlayerID `json:"playerID"`
	Field    string   `json:"field"`
	Value    string   `json:"value"`
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
	guessingPlayer.checks = append(guessingPlayer.checks, Check{Name: check.Field, GuessedValue: check.Value, Correct: check.Value == fact.Value})
	return nil
}

func (s *State) StartPregame() {
	s.Game.State = LOBBY
	// TODO Do we need or want this to do other state cleanup?
}

type StartGame struct {
	EventName
	StopTime *time.Time `json:"stopTime"`
}

func (s *State) StartGame(stopTime *time.Time) error {
	rand.Seed(time.Now().Unix())
	s.Game.State = PLAYING
	currentTime := time.Now()
	s.Game.StartTime = &currentTime
	s.Game.StopTime = stopTime
	if len(s.Facts.Real) == 0 {
		return errors.New("Cannot start a game with no real facts")
	}
	for name, fact := range s.Facts.Real {
		if len(fact.Possible) == 0 {
			return fmt.Errorf("fact %v has no possible values", name)
		}
		fact.Value = fact.Possible[rand.Intn(len(fact.Possible))]
	}
	return nil
}

type SubmitFacts struct {
	EventName
	Submission map[string]string `json:"submission"`
}

func (s *State) SubmitFacts(submission map[string]string) {
	s.Facts.Submitted = submission
}

type SendMessage struct {
	EventName
	GroupID GroupID `json:"groupID"`
	Sender  NodeID  `json:"sender"`
	Text    string  `json:"text"`
}

func (s *State) SendMessage(newMsg *SendMessage) error {
	group, ok := s.Graph.Groups[newMsg.GroupID]
	if !ok {
		return fmt.Errorf("group with id %v does not exist", newMsg.GroupID)
	}
	group.Messages = append(group.Messages, Message{Sender: newMsg.Sender, Timestamp: time.Now(), Text: newMsg.Text})
	return nil
}

type UpdateNode struct {
	EventName
	Node Node `json:"node"`
}

func (s *State) UpdateNode(newNode *UpdateNode) {
	s.Graph.Nodes[newNode.Node.ID] = &newNode.Node
}

type RemoveNode struct {
	EventName
	NodeID NodeID `json:"nodeID"`
}

func (s *State) RemoveNode(nodeID NodeID) {
	delete(s.Graph.Nodes, nodeID)
}

type UpdateGroup struct {
	EventName
	Group Group `json:"group"`
}

func (s *State) UpdateGroup(newGroup *UpdateGroup) {
	s.Graph.Groups[newGroup.Group.ID] = &newGroup.Group
}

type RemoveGroup struct {
	EventName
	GroupID GroupID `json:"groupID"`
}

func (s *State) RemoveGroup(groupID GroupID) {
	delete(s.Graph.Groups, groupID)
}

type UpdateRealFactPossibilities struct {
	EventName
	FactName       string   `json:"factName"`
	PossibleValues []string `json:"possibleValues"`
}

func (s *State) UpdateRealFactPossibilities(newFact *UpdateRealFactPossibilities) {
	updatedFact, ok := s.Facts.Real[newFact.FactName]
	if !ok {
		updatedFact = &Fact{}
		s.Facts.Real[newFact.FactName] = updatedFact
	}
	updatedFact.Possible = newFact.PossibleValues
}
