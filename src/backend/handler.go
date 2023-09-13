package backend

import (
	"html/template"
	"net/http"
)

<<<<<<< HEAD
=======
type TestGame struct {
	Tmp string
}

>>>>>>> ffa4ed30e3f00dac7bf73fd938476b02a9f8151a
func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).ParseFiles(filepaths...)
	// Error check:
	if err != nil {
		panic(err)
	}
	return tmpl
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
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
=======
	var game TestGame
	tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
	game = TestGame{
		Tmp: "test",
>>>>>>> ffa4ed30e3f00dac7bf73fd938476b02a9f8151a
	}
}
