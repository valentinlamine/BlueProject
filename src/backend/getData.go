package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// the array of unused events -> to make sure events appear only once
var notmade []int

// LoadEvents Function which loads the array of events from a json file
func LoadEvents(filename string) []Evt {
	f, _ := os.ReadFile(filename)
	var e []Evt
	err := json.Unmarshal(f, &e)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(e); i++ {
		notmade = append(notmade, i)
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
		fmt.Printf("Evenement -%v : %+v\n", i+1, e)
	}
}

// Remove Function which removes an element of an array
func Remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

// PickEvent Function which gives the next event, made is an array of id of past events
func PickEvent(events []Evt) Evt {
	if len(notmade) <= 0 {
		fmt.Println("Plus d'evenements")
		return Evt{}
	}
	i := rand.Intn(len(notmade))
	ind := notmade[i]
	notmade = Remove(notmade, i)
	return events[ind]
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
		fmt.Printf("Item -%v : %+v\n", i+1, e)
	}
}

// StartGame Function which initiates the data of the ntire game
func StartGame() Game {
	var g Game
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

// AddItem Function which adds the item from the index in the player inventory
func (game *Game) AddItem(ind int) {
	item := game.Items[ind]
	game.PlayerInfo.Inventory = append(game.PlayerInfo.Inventory, item)
}

// ApplyChoice upate player from the choice of the event
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

func (game *Game) ApplyResult(c Result) {
	game.PlayerInfo.Budget += c.Money
	game.PlayerInfo.Reputation += c.Reputation
	game.PlayerInfo.State += c.State
	if c.ObjectQuantity != 0 {
		game.AddItem(c.ObjectId)
	}
}
