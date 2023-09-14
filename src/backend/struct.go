package backend

// Game Original structure which contains all the others
type Game struct {
	PlayerInfo      Player
	AllEvents       []Evt
	FollowEvents    []Evt
	CurrentMarchant Marchant
	AllMarchants    []Marchant
	CurrentEvent    Evt
	Items           []Item
	FinalNotation   string
	MarchantTurn    int
	Turn            int
	Start           bool
	BonusReput      bool
}

// Item Struct Item
type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BuyPrice    int    `json:"buyPrice"`
	SellPrice   int    `json:"sellPrice"`
	Buyable     bool   `json:"achetable"`
	Money       int    `json:"money"`
	Reputation  int    `json:"reputation"`
	State       int    `json:"state"`
	Image       string `json:"image"`
}

// Result Struct Result, mainly
type Result struct {
	Money          int `json:"money"`
	Reputation     int `json:"reputation"`
	State          int `json:"state"`
	ObjectId       int `json:"object"`
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

type Marchant struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ItemsId     []int  `json:"itemsToSell"`
	Items       []Item
}
