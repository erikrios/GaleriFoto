package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/erikrios/cloud-native-programming-with-go/config"
	_ "github.com/erikrios/cloud-native-programming-with-go/config"
	"github.com/erikrios/cloud-native-programming-with-go/lib/persistence/mongodb"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	connection := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s/?authSource=admin", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection))
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
		}
	}()

	_ = mongodb.NewMongoDBLayer(client.Database(config.DBName))
	log.Println("Successfully connected into database")

	log.Println(fmt.Sprintf("Server started on port %d...\n", config.Port))
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

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}

type eventServiceHandler struct{}

func (e *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}
