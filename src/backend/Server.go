package backend

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func StartServer() {
// 	r := mux.NewRouter()
// 	r.PathPrefix("/ASSETS/").Handler(http.StripPrefix("/ASSETS/", http.FileServer(http.Dir("./ASSETS"))))
// 	r.PathPrefix("/CSS/").Handler(http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS"))))
// 	r.PathPrefix("/JS/").Handler(http.StripPrefix("/JS/", http.FileServer(http.Dir("./JS"))))

// 	//handle routing
// 	r.HandleFunc("/", IndexHandler)

// 	fmt.Println("server is running on port 8080 : http://localhost:8080")
// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		log.Fatal(err)
// 	}
// }
