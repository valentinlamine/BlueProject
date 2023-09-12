package backend

type Data struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	Descrition       string `json:"description"`
	LeftChoice       string `json:"leftChoice"`
	RightChoice      string `json:"rightChoice"`
	LeftChoiceResult struct {
		Money          int `json:"money"`
		Reputation     int `json:"reputation"`
		Etat           int `json:"etat"`
		ObjectId       int `json:"object-id"`
		ObjectQuantity int `json:"objectQuantity"`
	} `json:"leftChoiceResult"`
	RightChoiceResult struct {
		Money          int `json:"money"`
		Reputation     int `json:"reputation"`
		Etat           int `json:"etat"`
		ObjectId       int `json:"object-id"`
		ObjectQuantity int `json:"objectQuantity"`
	} `json:"rightChoiceResult"`
	LeftChoiceImage  string `json:"left-choice-image"`
	RightChoiceImage string `json:"right-choice-image"`
}
