package backend

import (
	"html/template"
	"net/http"
)

type TestGame struct {
	Tmp string
}

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).ParseFiles(filepaths...)
	// Error check:
	if err != nil {
		panic(err)
	}
	return tmpl
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var game TestGame
	tmpl := generateTemplate("index.html", []string{"frontend/index.html"})
	game = TestGame{
		Tmp: "test",
	}

	tmpl.Execute(w, game)
}
