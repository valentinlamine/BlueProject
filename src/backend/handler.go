package backend

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).ParseFiles(filepaths...)
	// Error check:
	if err != nil {
		panic(err)
	}
	return tmpl
}

func (g *Game) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(g.MarchantTurn, g.Turn)
	if r.Method == "POST" {
		if r.FormValue("name") != "" {
			fmt.Println("test1")
			g.PlayerInfo.Username = r.FormValue("name")
			r.Form.Set("name", "")
			tmpl := generateTemplate("game.html", []string{"frontend/game.html"})
			game := g.StartGame()
			g.SetupGame(r.FormValue("item"), r.FormValue("rep1"), r.FormValue("rep2"), r.FormValue("rep3"))
			g.Turn++
			tmpl.Execute(w, game)
		}
		if g.Turn%g.MarchantTurn == 0 {
			fmt.Println(g.Turn, "merchant turn")
			if r.FormValue("leave") != "" {
				fmt.Println("nextTurn")
				r.Form.Set("choice", "")
				tmpl := generateTemplate("game.html", []string{"frontend/game.html"})
				game := g.ContinueGame()
				g.Turn++
				tmpl.Execute(w, game)
			} else {
				fmt.Println("marchand")
				i := rand.Intn(3-0) + 0
				g.CurrentMarchant = g.AllMarchants[i]
				tmpl := generateTemplate("marchand.html", []string{"frontend/marchand.html"})
				tmpl.Execute(w, g)
			}
		} else {
			fmt.Println("nextTurn2")
			awnser, _ := strconv.Atoi(r.FormValue("choice"))
			g.ApplyChoice(awnser)
			r.Form.Set("choice", "")
			tmpl := generateTemplate("game.html", []string{"frontend/game.html"})
			game := g.ContinueGame()
			g.Turn++
			tmpl.Execute(w, game)
		}
	} else {
		tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
		tmpl.Execute(w, nil)
	}
}

func (g *Game) SetupGame(item string, rep1 string, rep2 string, rep3 string) {
	//change item to int
	itemInt, _ := strconv.Atoi(item)
	fmt.Println(itemInt)
	g.AddItem(itemInt)
	if rep1 == "left" {
		g.PlayerInfo.Reputation += 15
	} else {
		g.PlayerInfo.Reputation += 5
	}
	if rep2 == "left" {
		g.PlayerInfo.Reputation += 1
	} else {
		g.PlayerInfo.Reputation += 1
	}
	if rep3 == "left" {
		g.PlayerInfo.Reputation += 1
	} else {
		g.PlayerInfo.Reputation += 1
	}
}

func (g *Game) SellHandler(w http.ResponseWriter, r *http.Request) {
	/* response to api where user send the id of the item he wants to sell */
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the id of the item to sell
	var data struct {
		Id int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//inserer la fonction de vente ici
	success, info := g.SellItem(data.Id)
	if success {
		response := struct {
			Success    bool   `json:"success"`
			Info       string `json:"info"`
			Budget     int    `json:"budget"`
			Reputation int    `json:"reputation"`
			EtatEcole  int    `json:"etatEcole"`
		}{
			Success:    true,
			Info:       info,
			Budget:     g.PlayerInfo.Budget,
			Reputation: g.PlayerInfo.Reputation,
			EtatEcole:  g.PlayerInfo.State,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(jsonResponse)
	} else {
		response := struct {
			Success    bool   `json:"success"`
			Info       string `json:"info"`
			Budget     int    `json:"budget"`
			Reputation int    `json:"reputation"`
			EtatEcole  int    `json:"etatEcole"`
		}{
			Success:    false,
			Info:       info,
			Budget:     g.PlayerInfo.Budget,
			Reputation: g.PlayerInfo.Reputation,
			EtatEcole:  g.PlayerInfo.State,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(jsonResponse)
	}
}

func (g *Game) BuyHandler(w http.ResponseWriter, r *http.Request) {
	/* response to api where user send the id of the item he wants to buy */
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the id of the item to sell
	var data struct {
		Id int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//inserer la fonction d'achat ici
	success, info := g.BuyItem(data.Id)
	if success {
		response := struct {
			Success    bool   `json:"success"`
			Info       string `json:"info"`
			Budget     int    `json:"budget"`
			Reputation int    `json:"reputation"`
			EtatEcole  int    `json:"etatEcole"`
		}{
			Success:    true,
			Info:       info,
			Budget:     g.PlayerInfo.Budget,
			Reputation: g.PlayerInfo.Reputation,
			EtatEcole:  g.PlayerInfo.State,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(jsonResponse)
	} else {
		response := struct {
			Success    bool   `json:"success"`
			Info       string `json:"info"`
			Budget     int    `json:"budget"`
			Reputation int    `json:"reputation"`
			EtatEcole  int    `json:"etatEcole"`
		}{
			Success:    false,
			Info:       info,
			Budget:     g.PlayerInfo.Budget,
			Reputation: g.PlayerInfo.Reputation,
			EtatEcole:  g.PlayerInfo.State,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(jsonResponse)
	}
}

func (g *Game) UseHandler(w http.ResponseWriter, r *http.Request) {
	/* response to api where user send the id of the item he wants to use */
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the id of the item to sell
	var data struct {
		Id int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//inserer la fonction d'utilisation ici
	success, info := g.UseItem(data.Id)
	if success {
		response := struct {
			Success    bool   `json:"success"`
			Info       string `json:"info"`
			Budget     int    `json:"budget"`
			Reputation int    `json:"reputation"`
			EtatEcole  int    `json:"etatEcole"`
		}{
			Success:    true,
			Info:       info,
			Budget:     g.PlayerInfo.Budget,
			Reputation: g.PlayerInfo.Reputation,
			EtatEcole:  g.PlayerInfo.State,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(jsonResponse)
	} else {
		response := struct {
			Success    bool   `json:"success"`
			Info       string `json:"info"`
			Budget     int    `json:"budget"`
			Reputation int    `json:"reputation"`
			EtatEcole  int    `json:"etatEcole"`
		}{
			Success:    false,
			Info:       info,
			Budget:     g.PlayerInfo.Budget,
			Reputation: g.PlayerInfo.Reputation,
			EtatEcole:  g.PlayerInfo.State,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(jsonResponse)
	}
}
