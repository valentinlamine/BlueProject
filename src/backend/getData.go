package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// LoadEvents Function which loads the array of events from a json file
func LoadEvents(filename string) []Evt {
	f, _ := os.ReadFile(filename)
	var e []Evt
	err := json.Unmarshal(f, &e)
	if err != nil {
		log.Fatal(err)
	}

	return e
}

// PrintEvents Function which prints the array of events
func PrintEvents(events []Evt) {
	if len(events) == 0 {
		fmt.Println("Pas d'evenements ...")
		return
	}
	for i, e := range events {
		fmt.Println(e.Id, " ", i)
	}
}

// Remove Function which removes an element of an array
func Remove(slice []Evt, i int) []Evt {
	return append(slice[:i], slice[i+1:]...)
}

// LoadItems Function which loads the array of items from a json file
func LoadItems(filename string) []Item {
	f, _ := os.ReadFile(filename)
	var e []Item
	err := json.Unmarshal(f, &e)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

// PrintItems Function which prints the array of items
func PrintItems(events []Item) {
	if len(events) == 0 {
		fmt.Println("Pas d'items ...")
		return
	}
	for i, e := range events {
		fmt.Println(e.Id, " ", i)
	}
}

// StartGame Function which initiates the data of the ntire game
func StartGame(player Player) Game {
	var g Game
	g.PlayerInfo = player
	g.Items = LoadItems("DATA/items.json")
	g.AllEvents = EventShuffle(LoadEvents("DATA/events.json"))
	return g
}

// EventShuffle Function which randomizes the event array
func EventShuffle(events []Evt) []Evt {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(events),
		func(i, j int) { events[i], events[j] = events[j], events[i] })
	return events
}

func (game *Game) Following() {
	var fe []Evt
	var e []Evt = game.AllEvents

	fe = append(fe, e[21])
	fe = append(fe, e[9])
	fe = append(fe, e[4])
	fe = append(fe, e[2])

	Remove(e, 21)
	Remove(e, 9)
	Remove(e, 4)
	Remove(e, 2)
}

// AddItem Function which adds the item from the index in the player inventory
func (game *Game) AddItem(ind int) {
	item := game.Items[ind]
	game.PlayerInfo.Inventory = append(game.PlayerInfo.Inventory, item)
}

// ApplyChoice select the choice from an int
func (game *Game) ApplyChoice(choice int) {
	var c Result
	event := game.AllEvents[0]
	if choice == 0 {
		c = event.LeftResult
	} else if choice == 1 {
		c = event.RightResult
	} else {
		return
	}
	game.ApplyResult(c)
}

// ApplyResult update player from the choice of the event
func (game *Game) ApplyResult(c Result) {
	game.PlayerInfo.Budget += c.Money
	game.PlayerInfo.Reputation += c.Reputation
	game.PlayerInfo.State += c.State
	if c.ObjectQuantity != 0 {
		game.AddItem(c.ObjectId)
	}
}
