package routes

import (
	"github.com/gorilla/mux"
	"github.com/leonardogoandete/go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
