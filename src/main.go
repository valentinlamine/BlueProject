package main

import (
	"fmt"
	"log"
	"net/http"
	"src/backend"

	"github.com/gorilla/mux"
)

func main() {
	var g backend.Game
	r := mux.NewRouter()

	r.PathPrefix("/ASSETS/").Handler(http.StripPrefix("/ASSETS/", http.FileServer(http.Dir("./frontend/ASSETS"))))
	r.PathPrefix("/CSS/").Handler(http.StripPrefix("/CSS/", http.FileServer(http.Dir("./frontend/CSS"))))
	r.PathPrefix("/JS/").Handler(http.StripPrefix("/JS/", http.FileServer(http.Dir("./frontend/JS")))) //handle routing
	r.HandleFunc("/", g.IndexHandler)

	fmt.Println("server is running on port 8080 : http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
