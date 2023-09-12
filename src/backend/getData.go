package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetData() []Data {
	var data []Data
	jsonFile, err := os.Open("DATA/events.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &data)

	return data
}

// var dejaFait []int

// func RandomEvent() {

// 	var random int = rand.Intn(10)
// 	data := GetData()
// 	if len(dejaFait) == 0 {
// 		fmt.Println(data[random].Title)
// 		dejaFait = append(dejaFait, random)
// 	}

// 	for i := 0; i < len(dejaFait); i++ {
// 		if random == dejaFait[i] {
// 			RandomEvent()
// 		} else {
// 			fmt.Println(data[random].Title)
// 			fmt.Println(dejaFait)
// 			dejaFait = append(dejaFait, random)
// 		}
// 	}

// }
