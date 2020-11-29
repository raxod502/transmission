package model

import (
	"fmt"
	"math/rand"
	"strconv"
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
	rand.Seed(time.Now().UnixNano())
	possibleCompartments := []string{}
	for i := 0; i < 50; i++ {
		possibleCompartments = append(possibleCompartments, strconv.Itoa(i))
	}
	return &State{
		Game: Game{
			State: LOBBY,
		},
		Players: map[PlayerID]*Player{},
		Graph: Graph{
			Nodes: map[NodeID]*Node{
				"node-1": {
					ID:     "node-1",
					Groups: []GroupID{"group-1-2", "group-6-1", "group-1-4"},
				},
				"node-2": {
					ID:     "node-2",
					Groups: []GroupID{"group-1-2", "group-2-3", "group-2-5"},
				},
				"node-3": {
					ID:     "node-3",
					Groups: []GroupID{"group-2-3", "group-3-4", "group-3-6"},
				},
				"node-4": {
					ID:     "node-4",
					Groups: []GroupID{"group-3-4", "group-4-5", "group-1-4"},
				},
				"node-5": {
					ID:     "node-5",
					Groups: []GroupID{"group-4-5", "group-5-6", "group-2-5"},
				},
				"node-6": {
					ID:     "node-6",
					Groups: []GroupID{"group-5-6", "group-6-1", "group-3-6"},
				},
			},
			Groups: map[GroupID]*Group{
				"group-1-2": {
					ID:       "group-1-2",
					Messages: []Message{},
				},
				"group-2-3": {
					ID:       "group-2-3",
					Messages: []Message{},
				},
				"group-3-4": {
					ID:       "group-3-4",
					Messages: []Message{},
				},
				"group-4-5": {
					ID:       "group-4-5",
					Messages: []Message{},
				},
				"group-5-6": {
					ID:       "group-5-6",
					Messages: []Message{},
				},
				"group-6-1": {
					ID:       "group-6-1",
					Messages: []Message{},
				},
				"group-1-4": {
					ID:       "group-1-4",
					Messages: []Message{},
				},
				"group-2-5": {
					ID:       "group-2-5",
					Messages: []Message{},
				},
				"group-3-6": {
					ID:       "group-3-6",
					Messages: []Message{},
				},
				"group-baddies": {
					ID:       "group-baddies",
					Messages: []Message{},
				},
			},
		},
		Facts: Facts{
			Real:      map[string]*Fact{},
			Submitted: map[string]string{},
		},
		PossibleRoles: map[Role]int{
			HEADQUARTERS: 0,
			TRAINDEPOT:   1,
			DOUBLEAGENT:  2,
			TRAINEXPERT:  1,
			FACTCHECKER:  2,
		},
		PossibleFacts: map[string]*Fact{
			"compartment": {
				Possible: possibleCompartments,
				Value:    "42",
			},
			"color": {
				Possible: []string{
					"Black",
					"Blue",
					"Blue Green",
					"Blue Violet",
					"Brown",
					"Carnation Pink",
					"Dandelion",
					"Gray",
					"Green",
					"Cerulean",
					"Green Yellow",
					"Orange",
					"Apricot",
					"Scarlet",
					"Red",
					"Red Orange",
					"Red Violet",
					"Violet",
					"Indigo",
					"Violet Red",
					"White",
					"Yellow",
					"Yellow Green",
					"Yellow Orange",
				},
				Value: "violet red",
			},
			"food": {
				Possible: []string{
					"Roast Turkey",
					"Green Bean Casserole",
					"Candied Yams",
					"Mashed Potatoes",
					"Gravy",
					"Dry Brined Turkey",
					"Stuffing",
					"Cranberry Sauce",
					"Pumpkin Pie",
					"Cranberry Stuffing",
					"Ranch-seasoned Roast Turkey",
					"Cornbread Dressing",
					"Sausage Gravy",
					"Sweet Potato Casserole With Marshmallows",
					"Cornbread Stuffing",
					"Apple Pie",
					"Roasted Vegetables",
					"Turkey Cake",
					"Garlic Mashed Potatoes",
					"Roasted Butternut Squash Soup",
					"Brown Butter Mashed Potatoes",
					"Roasted Brussel Sprouts",
					"Mashed Sweet Potatoes",
					"Gravy Without Drippings",
					"Pecan Pie",
					"Sweet Potato Casserole",
					"Garlicky Green Beans with Crispy Onions",
					"Baked Sweet Potato",
					"Skillet Green Beans",
					"Cheesy Sweet Potato Casserole",
					"Butternut Squash and Andouille Stuffing",
					"Dark Chocolate Bourbon Pecan Pie",
					"Green Bean Casserole with Onion Rings",
				},
				Value: "pecan pie",
			},
		},
	}
}

type State struct {
	Game          Game                 `json:"game"`
	Players       map[PlayerID]*Player `json:"players"`
	Graph         Graph                `json:"graph"`
	Facts         Facts                `json:"facts"`
	PossibleRoles map[Role]int         `json:"possibleRoles"`
	PossibleFacts map[string]*Fact     `json:"possibleFacts"`
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
	Color      string          `json:"color"`
	Checks     []Check         `json:"checks"`
	KnownRoles map[NodeID]Role `json:"knownRoles"`
	KnownFacts map[string]bool `json:"knownFacts"`
	PowerUses  int             `json:"powerUses"`
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
	Groups []GroupID `json:"groups"` //TODO this should be a set
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

func (s *State) ResetGame() {
	// Delete messages from current game
	for _, group := range s.Graph.Groups {
		group.Messages = []Message{}
	}
	// Keep only id, name, and color for player
	for id, player := range s.Players {
		s.Players[id] = &Player{
			Name:  player.Name,
			Color: player.Color,
			ID:    player.ID,
			Node:  "",
		}
	}
	// Reset facts
	s.Facts.Real = map[string]*Fact{}
	s.Facts.Submitted = map[string]string{}
	// Clear node assignments but keep map structure
	for id, node := range s.Graph.Nodes {
		s.Graph.Nodes[id] = &Node{
			ID:     node.ID,
			Groups: node.Groups,
		}
	}

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

func (s *State) StopGame() error {
	s.Game.State = SUBMISSION
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
	guessingPlayer.Checks = append(guessingPlayer.Checks, Check{Name: check.Field, GuessedValue: check.Value, Correct: check.Value == fact.Value})
	guessingPlayer.PowerUses++
	return nil
}

func (s *State) StartPregame() {
	s.Game.State = LOBBY
	s.ResetGame()
}

type StartGame struct {
	EventName
	StopTime *time.Time `json:"stopTime"`
}

func (s *State) StartGame(stopTime *time.Time) error {
	if len(s.Players) != len(s.Graph.Nodes) {
		return fmt.Errorf("cannot start a game on a map with %v nodes but only %v player(s)", len(s.Graph.Nodes), len(s.Players))
	}
	// set start and stop time
	rand.Seed(time.Now().Unix())
	s.Game.State = PLAYING
	currentTime := time.Now()
	s.Game.StartTime = &currentTime
	s.Game.StopTime = stopTime
	// Default to using all facts
	if len(s.Facts.Real) == 0 {
		s.Facts.Real = s.PossibleFacts
	}
	// Randomly select fact values
	for name, fact := range s.Facts.Real {
		if len(fact.Possible) == 0 {
			return fmt.Errorf("fact %v has no possible values", name)
		}
		fact.Value = fact.Possible[rand.Intn(len(fact.Possible))]
	}
	// Assign roles and nodes to any players who don't already have them
	unassignedPlayers := []*Player{}
	playersWithoutRoles := []*Player{}
	usedRoles := map[Role]int{}
	for _, player := range s.Players {
		if player.Node == "" {
			unassignedPlayers = append(unassignedPlayers, player)
		}
		if player.Role == "" {
			playersWithoutRoles = append(playersWithoutRoles, player)
		} else {
			usedRoles[player.Role]++
		}
	}
	if len(unassignedPlayers) != 0 {
		err := s.autoAssignPlayersToNodes(unassignedPlayers)
		if err != nil {
			return err
		}
	}
	if len(playersWithoutRoles) != 0 {
		err := s.autoAssignPlayersRoles(playersWithoutRoles, usedRoles)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *State) autoAssignPlayersToNodes(unassignedPlayers []*Player) error {
	emptyNodes := []*Node{}
	for _, node := range s.Graph.Nodes {
		if node.Player == "" {
			emptyNodes = append(emptyNodes, node)
		}
	}
	if len(emptyNodes) != len(unassignedPlayers) {
		return fmt.Errorf("%v empty nodes but %v unassigned players", len(emptyNodes), len(unassignedPlayers))
	}
	rand.Shuffle(len(emptyNodes), func(i int, j int) { emptyNodes[i], emptyNodes[j] = emptyNodes[j], emptyNodes[i] })
	for i, player := range unassignedPlayers {
		player.Node = emptyNodes[i].ID
		emptyNodes[i].Player = player.ID
		emptyNodes[i].Color = player.Color
		emptyNodes[i].Name = player.Name
	}
	return nil
}

func (s *State) autoAssignPlayersRoles(playersWithoutRoles []*Player, usedRoles map[Role]int) error {
	rolesLeft := map[Role]int{}
	for role, possibleCount := range s.PossibleRoles {
		rolesLeft[role] = possibleCount - usedRoles[role]
	}
	rolesToAssign := []Role{}
	for role, countLeft := range rolesLeft {
		if countLeft < 0 {
			return fmt.Errorf("Cannot automatically assign roles because there are too many players with the role %v", role)
		}
		for i := 0; i < countLeft; i++ {
			rolesToAssign = append(rolesToAssign, role)
		}
	}
	if len(rolesToAssign) != len(playersWithoutRoles) {
		return fmt.Errorf("%v players without roles, but only %v roles to assign", len(playersWithoutRoles), len(rolesToAssign))
	}
	rand.Shuffle(len(rolesToAssign), func(i, j int) { rolesToAssign[i], rolesToAssign[j] = rolesToAssign[j], rolesToAssign[i] })
	for i, player := range playersWithoutRoles {
		player.Role = rolesToAssign[i]
	}
	return nil
}

type SubmitFacts struct {
	EventName
	Submission map[string]string `json:"submission"`
}

func (s *State) SubmitFacts(submission map[string]string) {
	s.Facts.Submitted = submission
	s.Game.State = RESULTS
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

type SetRealFacts struct {
	EventName
	Facts map[string]*Fact
}

func (s *State) SetRealFacts(facts *SetRealFacts) {
	s.Facts.Real = facts.Facts
}

type AddKnownRole struct {
	EventName
	PlayerID PlayerID
	NodeID   NodeID
	Role     Role
}

func (s *State) AddKnownRole(message *AddKnownRole) error {
	player, ok := s.Players[message.PlayerID]
	if !ok {
		return fmt.Errorf("player with id %v does not exist", message.PlayerID)
	}
	if player.KnownRoles == nil {
		player.KnownRoles = map[NodeID]Role{}
	}
	player.KnownRoles[message.NodeID] = message.Role
	player.PowerUses++
	return nil
}

type AddKnownFact struct {
	EventName
	PlayerID PlayerID
	Names    []string
}

func (s *State) AddKnownFact(message *AddKnownFact) error {
	player, ok := s.Players[message.PlayerID]
	if !ok {
		return fmt.Errorf("player with id %v does not exist", message.PlayerID)
	}
	if player.KnownFacts == nil {
		player.KnownFacts = map[string]bool{}
	}
	for _, name := range message.Names {
		player.KnownFacts[name] = true
	}
	return nil
}

type UpdatePossibleRoles struct {
	EventName
	PossibleRoles map[Role]int
}

func (s *State) UpdatePossibleRoles(message *UpdatePossibleRoles) error {
	s.PossibleRoles = message.PossibleRoles
	return nil
}
