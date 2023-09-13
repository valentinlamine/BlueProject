package backend

import (
	"fmt"
	"html/template"
	"net/http"
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
	var p Player
	if p.Username != "" {
		fmt.Println("gné")
		tmpl := generateTemplate("game.html", []string{"frontend/game.html"})
		game := StartGame(p)
		tmpl.Execute(w, game)
	} else if r.Method == "POST" {
		if r.FormValue("name") != "" {
			p.Username = r.FormValue("name")
		} else {
			tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
			tmpl.Execute(w, nil)
		}
	} else {
		tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
		tmpl.Execute(w, nil)
	}
}
