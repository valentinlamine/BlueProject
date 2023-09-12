package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	r.PathPrefix("/CSS/").Handler(http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS"))))
	r.PathPrefix("/JS/").Handler(http.StripPrefix("/JS/", http.FileServer(http.Dir("./JS"))))

	//handle routing
	r.HandleFunc("/", utilities.IndexHandler)
	r.HandleFunc("/following", utilities.FollowedHandler)
	r.HandleFunc("/admin", utilities.AdminHandler)
	r.HandleFunc("/topic/{id}", utilities.TopicHandler)
	r.HandleFunc("/termsofservice", utilities.TermsOfServiceHandler)
	r.HandleFunc("/privacypolicy", utilities.PrivacyPolicyHandler)
	r.HandleFunc("/test", utilities.TestHandler)

	r.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World2") }).Methods("GET")

	r.HandleFunc("/login", utilities.LoginHandler).Methods("POST")
	r.HandleFunc("/register", utilities.RegisterHandler).Methods("POST")
	r.HandleFunc("/topicCreate", utilities.TopicCreateHandler).Methods("POST")
	r.HandleFunc("/topicDelete", utilities.TopicDeleteHandler).Methods("POST")
	r.HandleFunc("/messageCreate", utilities.MessageCreateHandler).Methods("POST")
	r.HandleFunc("/messageDelete", utilities.MessageDeleteHandler).Methods("POST")
	//r.HandleFunc("/messageEdit", utilities.MessageEditHandler).Methods("POST")
	r.HandleFunc("/messageInteractions", utilities.MessageInteractionsHandler).Methods("POST")

	r.HandleFunc("/bookmarkAdd", utilities.BookmarkAddHandler).Methods("POST")
	r.HandleFunc("/bookmarkRemove", utilities.BookmarkRemoveHandler).Methods("POST")
	r.HandleFunc("/addRoleHandler", utilities.AddRoleHandler).Methods("POST")
	r.HandleFunc("/SupprUserRoleHandler", utilities.SupprUserRoleHandler).Methods("POST")
	r.HandleFunc("/createRoleHandler", utilities.CreateRoleHandler).Methods("POST")
	r.HandleFunc("/SupprRoleHandler", utilities.SupprRoleHandler).Methods("POST")

	fmt.Println("server is running on port 8080 : http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
