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
func (g *Game) StartGame(player Player) Game {
	g.PlayerInfo = player
	g.Items = LoadItems("DATA/items.json")
	g.AllEvents = LoadEvents("DATA/events.json")
	g.Following()
	g.AllEvents = EventShuffle(g.AllEvents)
	g.CurrentEvent = g.AllEvents[0]
	return *g
}

func (g *Game) ContinueGame() Game {
	g.AllEvents = g.AllEvents[1:]
	g.CurrentEvent = g.AllEvents[0]
	return *g
}

// EventShuffle Function which randomizes the event array
func EventShuffle(events []Evt) []Evt {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(events),
		func(i, j int) { events[i], events[j] = events[j], events[i] })
	return events
}

// Remove Function which removes an element of an array
func Remove(slice []Evt, i int) []Evt {
	return append(slice[:i], slice[i+1:]...)
}

// Remove Function which removes an element of an array
func RemoveItem(slice []Item, i int) []Item {
	return append(slice[:i], slice[i+1:]...)
}

// Following separates events with conditons and normal events
func (g *Game) Following() {
	g.FollowEvents = append(g.FollowEvents, g.AllEvents[21], g.AllEvents[9], g.AllEvents[4], g.AllEvents[2])
	g.AllEvents = Remove(g.AllEvents, 21)
	g.AllEvents = Remove(g.AllEvents, 9)
	g.AllEvents = Remove(g.AllEvents, 4)
	g.AllEvents = Remove(g.AllEvents, 2)

	fmt.Println("\n----------------")
	for i := 0; i != len(g.AllEvents); i++ {
		fmt.Print("id:", g.AllEvents[i].Id, " ")
	}
}

// AddItem Function which adds the item from the index in the player inventory
func (game *Game) AddItem(ind int) {
	item := game.Items[ind]
	game.PlayerInfo.Inventory = append(game.PlayerInfo.Inventory, item)
}

// BuyItem return 1 if buying is impossible an 0 if it is possible
func (game *Game) BuyItem(ind int) int {
	item := game.Items[ind]
	if item.BuyPrice > game.PlayerInfo.Budget {
		return 1
	}
	game.AddItem(ind)
	game.PlayerInfo.Budget -= item.BuyPrice
	return 0
}

// SellItem
func (game *Game) SellItem(ind int) {
	item := game.Items[ind]
	game.PlayerInfo.Budget += item.SellPrice
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

// AddEventadds the follow event
func (game *Game) AddEvent(id int, choice int) Evt {
	var event Evt

	return event
}
