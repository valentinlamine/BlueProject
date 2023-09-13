package backend

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	if r.Method == "POST" {
		if r.FormValue("name") != "" {
			g.PlayerInfo.Username = r.FormValue("name")
			//clear form value
			r.Form.Set("name", "")
			tmpl := generateTemplate("game.html", []string{"frontend/game.html"})
			game := g.StartGame()
			fmt.Println(r.FormValue("choice"))
			tmpl.Execute(w, game)
		}
		if r.FormValue("choice") != "" {
			awnser, _ := strconv.Atoi(r.FormValue("choice"))
			g.ApplyChoice(awnser)
			//r.FormValue("choice") = ""

			tmpl := generateTemplate("game.html", []string{"frontend/game.html"})
			game := g.ContinueGame()
			fmt.Println(r.FormValue("choice"))
			tmpl.Execute(w, game)
		}
	} else {
		tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
		tmpl.Execute(w, nil)
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
	success, info := true, "Item vendu"
	if success {
		response := struct {
			Success bool   `json:"success"`
			Info    string `json:"info"`
		}{
			Success: true,
			Info:    info,
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
			Success bool   `json:"success"`
			Info    string `json:"info"`
		}{
			Success: false,
			Info:    info,
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
	success, info := true, "Item acheté"
	if success {
		response := struct {
			Success bool   `json:"success"`
			Info    string `json:"info"`
		}{
			Success: true,
			Info:    info,
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
			Success bool   `json:"success"`
			Info    string `json:"info"`
		}{
			Success: false,
			Info:    info,
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
