package backend

import (
	"encoding/json"
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

func LoadMarchand(filename string, g Game) []Marchant {
	f, _ := os.ReadFile(filename)
	var e []Marchant
	err := json.Unmarshal(f, &e)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i != len(e); i++ {
		for j := 0; j != len(e[i].ItemsId); j++ {
			e[i].Items = append(e[i].Items, g.GetItemById(e[i].ItemsId[j]))
		}
	}
	return e
}

// StartGame Function which initiates the data of the entire game
func (g *Game) StartGame() Game {
	g.Turn = 0
	g.PlayerInfo.Reputation = 0
	g.PlayerInfo.Budget = 6700
	g.PlayerInfo.State = 50
	g.Items = LoadItems("DATA/items.json")
	g.AllMarchants = LoadMarchand("DATA/trader.json", *g)
	g.AllEvents = LoadEvents("DATA/events.json")
	g.MarchantTurn = 2 + (len(g.AllEvents) / 3)
	g.Following()
	g.EventShuffle(g.AllEvents)
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
func (g *Game) EventShuffle(events []Evt) {
	//var tmp Evt
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(events),
		func(i, j int) { events[i], events[j] = events[j], events[i] })
	/*for i := 1; i != len(events); i++ {
		if i%g.MarchantTurn == 0 {
			Insert(g.AllEvents, i, tmp)
		}
	}*/
	g.AllEvents = events
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
	g.FollowEvents = append(g.FollowEvents, g.AllEvents[20], g.AllEvents[9], g.AllEvents[4], g.AllEvents[2])
	g.AllEvents = Remove(g.AllEvents, 20)
	g.AllEvents = Remove(g.AllEvents, 9)
	g.AllEvents = Remove(g.AllEvents, 4)
	g.AllEvents = Remove(g.AllEvents, 2)
}

// AddItem Function which adds the item from the index in the player inventory
func (game *Game) AddItem(id int) {
	item := game.Items[id-1]
	game.PlayerInfo.Inventory = append(game.PlayerInfo.Inventory, item)
}

// BuyItem return 1 if buying is impossible an 0 if it is possible
func (game *Game) BuyItem(id int) (bool, string) {
	for i := 0; i < len(game.PlayerInfo.Inventory); i++ {
		if game.PlayerInfo.Inventory[i].Id == id {
			return false, "Objet déjà possédé"
		}
	}
	item := game.Items[id-1]
	if item.BuyPrice > game.PlayerInfo.Budget {
		return false, "Pas assez d'argent"
	}
	game.AddItem(id)
	game.PlayerInfo.Budget -= item.BuyPrice
	return true, "item acheté"
}

// SellItem Removes the item and adds the money to the player
func (game *Game) SellItem(id int) (bool, string) {
	var ind int
	var b bool = false
	for i := 0; i < len(game.PlayerInfo.Inventory); i++ {
		if game.PlayerInfo.Inventory[i].Id == id {
			b = true
			ind = i
		}
	}
	if !b {
		return false, "Item non possédé"
	}
	item := game.Items[id-1]
	game.PlayerInfo.Budget += item.SellPrice
	game.PlayerInfo.Inventory = RemoveItem(game.PlayerInfo.Inventory, ind)
	return true, "item vendu"
}

// ApplyChoice select the choice from an int
func (game *Game) ApplyChoice(choice int) (bool, string) {
	// event trahison
	if game.CurrentEvent.Id == 21 {
		var b bool = false
		for i := 0; i < len(game.PlayerInfo.Inventory); i++ {
			if game.PlayerInfo.Inventory[i].Id == 9 {
				b = true
			}
		}
		if b && choice == 1 {
			return true, "tout va bien"
		}
		return false, "Prison"
	}
	// event incendie
	if game.CurrentEvent.Id == 10 {
		if choice == 1 {
			return false, "Incendie"
		}
		if game.PlayerInfo.Budget <= 15000 {
			return false, "Incendie"
		}
	}

	var c Result
	event := game.CurrentEvent
	if choice == 0 {
		c = event.LeftResult
	} else if choice == 1 {
		c = event.RightResult
	} else {
		return false, "Erreur"
	}
	return game.ApplyResult(c)
}

// ApplyResult update player from the choice of the event
func (game *Game) ApplyResult(c Result) (bool, string) {
	game.PlayerInfo.Budget += c.Money
	if game.PlayerInfo.Budget <= 0 {
		var ind int
		var b bool = false
		for i := 0; i < len(game.PlayerInfo.Inventory); i++ {
			if game.PlayerInfo.Inventory[i].Id == 4 {
				ind = i
				game.PlayerInfo.Inventory = RemoveItem(game.PlayerInfo.Inventory, ind)
				b = true
			}
		}
		if !b {
			return false, "Banqueroute"
		}
	}

	// reputation can must be between 100 and -100
	game.PlayerInfo.Reputation += c.Reputation
	if game.PlayerInfo.Reputation > 100 {
		game.PlayerInfo.Reputation = 100
	}
	if game.PlayerInfo.Reputation < -100 {
		game.PlayerInfo.Reputation = -100
	}

	// state must be between 0 and 100
	game.PlayerInfo.State += c.State
	if game.PlayerInfo.State <= 0 {
		return false, "Ecroulement"
	}
	if game.PlayerInfo.State > 100 {
		game.PlayerInfo.State = 100
	}

	// add the object if necessary
	if c.ObjectQuantity != 0 {
		game.AddItem(c.ObjectId)
	}
	return true, "tout va bien"
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
func (game *Game) ManageEvent(choice int) (bool, string) {
	id := game.CurrentEvent.Id
	ok, s := game.ApplyChoice(choice)
	if !ok {
		return ok, s
	}
	if len(game.AllEvents) == 1 {
		return false, "Victoire"
	}

	switch id {
	case 2:
		if choice == 1 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[3])
		}
	case 4:
		if choice == 1 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[2])
		}
	case 9:
		if choice == 0 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[1])
		}
	case 19:
		if choice == 1 {
			game.AllEvents = Insert(game.AllEvents, 1, game.FollowEvents[0])
		}
	}

	return true, "Tout va bien"
}

// UseItem Triggers the item effect and destroy it
func (game *Game) UseItem(id int) (bool, string) {
	// ignoring the specials items
	if id == 4 || id == 9 {
		return false, "Item non consommable"
	}

	// skipping the current event with the item
	if id == 8 {
		game.ContinueGame()
	}

	// consumables
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

	return true, "item utilisé"
}

func (game *Game) GetItemById(id int) Item {
	for _, item := range game.Items {
		if item.Id == id {
			return item
		}
	}
	return Item{}
}

func (g *Game) GetFinalNotation() string {
	switch true {
	case (g.PlayerInfo.Reputation >= -100 && g.PlayerInfo.Reputation <= -21):
		return "Vous êtes un démon"
	case (g.PlayerInfo.Reputation >= -20 && g.PlayerInfo.Reputation <= 0):
		return "Vous êtes mal"
	case (g.PlayerInfo.Reputation >= 1 && g.PlayerInfo.Reputation <= 20):
		return "Vous êtes bon"
	case (g.PlayerInfo.Reputation >= 21 && g.PlayerInfo.Reputation <= 100):
		return "Vous êtes un ange"
	default:
		return "vous êtres un ange"
	}
}
