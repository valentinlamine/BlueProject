package backend

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	for _, e := range events {
		fmt.Println(e.Id)
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
	for _, e := range events {
		fmt.Print(e.Id)
	}
}

// StartGame Function which initiates the data of the entire game
func (g *Game) StartGame() Game {
	g.PlayerInfo.Reputation = 0
	g.PlayerInfo.Budget = 20000
	g.PlayerInfo.State = 50
	g.Items = LoadItems("DATA/items.json")
	g.AllEvents = LoadEvents("DATA/events.json")
	g.Following()
	g.AllEvents = EventShuffle(g.AllEvents)
	g.CurrentEvent = g.AllEvents[0]
	g.Start = true
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

// RemoveItem Function which removes an element of an array
func RemoveItem(slice []Item, i int) []Item {
	return append(slice[:i], slice[i+1:]...)
}

// Following separates events with conditions and normal events
func (g *Game) Following() {
	g.FollowEvents = append(g.FollowEvents, g.AllEvents[21], g.AllEvents[9], g.AllEvents[4], g.AllEvents[2])
	g.AllEvents = Remove(g.AllEvents, 20)
	g.AllEvents = Remove(g.AllEvents, 9)
	g.AllEvents = Remove(g.AllEvents, 4)
	g.AllEvents = Remove(g.AllEvents, 2)

	/*fmt.Println("\n----------------")
	for i := 0; i != len(g.AllEvents); i++ {
		fmt.Print("id:", g.AllEvents[i].Id, " ")
	}*/
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

// SellItem Removes the item and adds the money to the player
func (game *Game) SellItem(ind int) {
	item := game.Items[ind]
	game.PlayerInfo.Budget += item.SellPrice
	game.PlayerInfo.Inventory = RemoveItem(game.PlayerInfo.Inventory, ind)
}

// ApplyChoice select the choice from an int
func (game *Game) ApplyChoice(choice int) int {
	var c Result
	event := game.CurrentEvent
	if choice == 0 {
		c = event.LeftResult
	} else if choice == 1 {
		c = event.RightResult
	} else {
		return 1
	}
	return game.ApplyResult(c)
}

// ApplyResult update player from the choice of the event
func (game *Game) ApplyResult(c Result) int {
	game.PlayerInfo.Budget += c.Money
	if game.PlayerInfo.Budget <= 0 {
		var ind int
		for i := 0; i < len(game.PlayerInfo.Inventory); i++ {
			if game.PlayerInfo.Inventory[i].Id == 4 {
				ind = i
				game.PlayerInfo.Inventory = RemoveItem(game.PlayerInfo.Inventory, ind)
			}
		}
		return 1
	}
	if game.PlayerInfo.Budget > 100 {
		game.PlayerInfo.Budget = 100
	}
	game.PlayerInfo.Reputation += c.Reputation
	game.PlayerInfo.State += c.State
	if game.PlayerInfo.State <= 0 {
		return 1
	}
	if game.PlayerInfo.State > 100 {
		game.PlayerInfo.State = 100
	}
	if c.ObjectQuantity != 0 {
		game.AddItem(c.ObjectId)
	}
	return 0
}

// Insert Allows to insert an element in an array
func Insert(a []Evt, index int, value Evt) []Evt {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

// ManageEvent adds the follow event and manage the current event, if the player looses, return 1 else 0
func (game *Game) ManageEvent(choice int) int {
	id := game.CurrentEvent.Id
	if game.ApplyChoice(choice) == 1 {
		return 1
	}

	switch id {
	case 2:
		if choice == 1 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[3])
			//fmt.Println("voici l'evenement choc: ", game.FollowEvents[3])
			//game.AllEvents = Insert(game.AllEvents, ind, game.FollowEvents[3])
		}
	case 4:
		if choice == 1 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[2])
			//fmt.Println("voici l'evenement choc: ", game.FollowEvents[2])
			//game.AllEvents = Insert(game.AllEvents, ind, game.FollowEvents[2])
		}
	case 9:
		if choice == 0 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[1])
			//fmt.Println("voici l'evenement choc: ", game.FollowEvents[1])
			//game.AllEvents = Insert(game.AllEvents, ind, game.FollowEvents[1])
		}
	case 19:
		if choice == 1 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[0])
			//fmt.Println("voici l'evenement choc: ", game.FollowEvents[0])
			//game.AllEvents = Insert(game.AllEvents, ind, game.FollowEvents[0])
		}
	}

	return 0
}

// UseItem Triggers the item effect and destroy it
func (game *Game) UseItem(id int) int {
	if id == 4 || id == 8 || id == 9 {
		return 1
	}
	ind := 0
	for i := 0; i < len(game.PlayerInfo.Inventory); i++ {
		if game.PlayerInfo.Inventory[i].Id == id {
			ind = 0
		}
	}
	item := game.PlayerInfo.Inventory[ind]
	game.PlayerInfo.Budget += item.Money
	game.PlayerInfo.Reputation += item.Reputation
	game.PlayerInfo.State += item.State

	game.PlayerInfo.Inventory = RemoveItem(game.PlayerInfo.Inventory, ind)

	return 0
}

// Test gathers the functions in order to have a playable game in terminal
func Test(player Player) {
	var game Game
	game.StartGame()
	PrintEvents(game.AllEvents)
	fmt.Println("len de events: ", len(game.AllEvents))

	// iteration until end of event array
	for ind := 0; ind < len(game.AllEvents); ind++ {
		if ind == len(game.AllEvents) {
			fmt.Println("Victoire")
			return
		}
		game.CurrentEvent = game.AllEvents[ind]

		// setup of interface
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("buget: ", game.PlayerInfo.Budget, " etat: ", game.PlayerInfo.State, " reput: ", game.PlayerInfo.Reputation)
		fmt.Println("-------------------------------------")
		fmt.Println("id: ", game.CurrentEvent.Id, "titre: ", game.CurrentEvent.Title)
		fmt.Println(game.CurrentEvent.Description)
		fmt.Println("choix 0: ", game.CurrentEvent.LeftChoice)
		fmt.Println("choix 1: ", game.CurrentEvent.RightChoice)
		for {
			fmt.Println("Enter choice: ")
			res, _ := reader.ReadString('\n')
			res = strings.Replace(res, "\n", "", -1)
			res = strings.Replace(res, "\r", "", -1)

			// getting the choice of user and managing the event
			choice, err := strconv.Atoi(res)
			fmt.Println("entered number: ", choice, "\nerreur: ", err)
			destin := game.ManageEvent(choice)
			if destin == 1 {
				fmt.Println("Defaite")
				return
			}
			break
		}
	}
}
