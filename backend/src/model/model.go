package model

import (
	"fmt"
	"math/rand"
	"time"
)

type Role string

const (
	HQ  Role = "HQ"
	TD  Role = "TD"
	BAD Role = "Baddie"
	TE  Role = "Train Expert"
	FC  Role = "Fact Checker"
)

type PlayerID int
type Fact struct {
	Value          string
	PossibleValues []string
}
type Facts map[string]Fact

type Thread []Message

type Message struct {
	Sender    *Player
	Text      string
	Timestamp time.Time
}

type Player struct {
	Name    string
	Role    Role
	ID      PlayerID
	Color   int // TODO: figure out best way to encode this
	Threads []*Thread
}

type Game struct {
	Players []*Player
	Threads map[PlayerID]map[PlayerID]*Thread
	Graph   map[PlayerID][]PlayerID
}

type RoleConfig struct {
	GoodRoles  []Role
	BadRoles   []Role
	NumBaddies int
}

func shuffleRoleSlice(slice []Role) {
	rand.Shuffle(
		len(slice),
		func(i int, j int) {
			slice[j], slice[i] = slice[i], slice[j]
		},
	)
}

func (rc *RoleConfig) randomizeRoles(numPlayers int) []Role {
	// Copy both lists of roles
	goods := make([]Role, len(rc.GoodRoles))
	copy(goods, rc.GoodRoles)
	bads := make([]Role, len(rc.BadRoles))
	copy(bads, rc.BadRoles)
	// Shuffle both lists
	shuffleRoleSlice(goods)
	shuffleRoleSlice(bads)
	// Take rc.NumBaddies bad roles
	chosenBadRoles := bads[:rc.NumBaddies]
	// Take numPlayers - rc.NumBaddies good roles
	chosenGoodRoles := goods[:(numPlayers - rc.NumBaddies)]
	// stick em together and shuffle them one more time
	chosenRoles := append(chosenBadRoles, chosenGoodRoles...)
	shuffleRoleSlice(chosenRoles)
	return chosenRoles
}

// Takes in a graph in adjacency list form
func MakeGame(graph map[int][]int, roles RoleConfig) Game {
	rand.Seed(time.Now().UTC().UnixNano())
	var g Game
	g.Threads = map[PlayerID]map[PlayerID]*Thread{}
	g.Graph = map[PlayerID][]PlayerID{}
	nodeToPlayer := map[int]*Player{}

	// Create a player for each node in the graph
	currentPlayerID := PlayerID(0)
	for node := range graph {
		player := &Player{ID: currentPlayerID, Name: fmt.Sprintf("Player %v", currentPlayerID)}
		currentPlayerID++

		nodeToPlayer[node] = player
		g.Players = append(g.Players, player)
	}

	// Handle edges
	for nodeA, neighbors := range graph {
		for _, nodeB := range neighbors {
			if nodeA > nodeB {
				// create a thread once for each edge in the graph
				thread := &Thread{}
				// Add thread to both players
				playerA := nodeToPlayer[nodeA]
				playerA.Threads = append(playerA.Threads, thread)
				playerB := nodeToPlayer[nodeB]
				playerB.Threads = append(playerB.Threads, thread)

				// Add thread to the game's thread mapping
				secondPlayerMap, ok := g.Threads[playerA.ID]
				if !ok {
					// initialize map if it doesn't already exist
					g.Threads[playerA.ID] = map[PlayerID]*Thread{}
					secondPlayerMap = g.Threads[playerA.ID]
				}
				secondPlayerMap[playerB.ID] = thread

				// Add the edge using the correct player IDs to game.Graph
				g.Graph[playerA.ID] = append(g.Graph[playerA.ID], playerB.ID)
			}
		}
	}
	// Assign roles
	randomRoles := roles.randomizeRoles(len(graph))
	// Player N gets role N (since role ordering is already random)
	for i, player := range g.Players {
		player.Role = randomRoles[i]
	}
	return g
}
