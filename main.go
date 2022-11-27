package main

import (
	"log"
	"net/http"

	_ "github.com/erikrios/cloud-native-programming-with-go/config"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Server started on port 8181...")
	if err := ServeAPI(); err != nil {
		log.Fatalln(err)
	}
}

func ServeAPI() error {
	handler := &eventServiceHandler{}

	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()

	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)

	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(":8181", r)
}

type eventServiceHandler struct{}

func (e *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}
