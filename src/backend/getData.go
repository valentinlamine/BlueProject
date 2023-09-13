package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// Game Original structure which contains all the others
type Game struct {
	PlayerInfo Player
	AllEvents  []Evt
	Items      []Item
}

// Item Struct Item
type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BuyPrice    int    `json:"buyPrice"`
	SellPrice   int    `json:"sellPrice"`
	Effect      int    `json:"effects"`
}

// Result Struct Result, mainly
type Result struct {
	Money          int `json:"money"`
	Reputation     int `json:"reputation"`
	Stat           int `json:"etat"`
	ObjectId       int `json:"object-id"`
	ObjectQuantity int `json:"ObjectQuantity"`
}

// Evt Struct event, which needs to be loaded from json
type Evt struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LeftChoice  string `json:"left-choice"`
	RightChoice string `json:"right-choice"`
	LeftResult  Result `json:"left-choice-result"`
	RightResult Result `json:"right-choice-result"`
	LeftImage   string `json:"left-choice-image"`
	RightImage  string `json:"right-choice-image"`
}

// Player Struct player, which needs to be loaded from json
type Player struct {
	Username   string
	Reputation int
	Budget     int
	Inventory  []Item
}

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
	notmade = Remove(notmade, i)
	return events[i]
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
