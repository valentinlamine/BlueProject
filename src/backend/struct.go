package backend

// Game Original structure which contains all the others
type Game struct {
	PlayerInfo Player
	AllEvents  []Evt
	Items      []Item
}

// Item Struct Item
type Item struct {
	Id          int    `json:"id"`
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
	State          int `json:"state"`
	ObjectId       int `json:"object-id"`
	ObjectQuantity int `json:"ObjectQuantity"`
}

// Evt Struct event, which needs to be loaded from json
type Evt struct {
	Id          int    `json:"id"`
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
	State      int
	Budget     int
	Inventory  []Item
}
