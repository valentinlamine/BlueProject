package backend

import (
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
	if r.Method == "POST" {
		if r.FormValue("start") == "start" {
			if r.FormValue("name") != "" {
				tmpl := generateTemplate("game.html", []string{"frontend/game.html"})
				game := StartGame()
				tmpl.Execute(w, game)
			} else {
				tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
				tmpl.Execute(w, nil)
			}
		}
	} else {
		tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
		tmpl.Execute(w, nil)
	}
}
