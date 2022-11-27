package mongodb

import (
	"github.com/erikrios/cloud-native-programming-with-go/lib/persistence"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	EVENTS = "events"
)

type MongoDBLayer struct {
	db *mongo.Database
}

func NewMongoDBLayer(db *mongo.Database) persistence.DatabaseHandler {
	return &MongoDBLayer{
		db: db,
	}
}

func (m *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	// s := m.getFreshSession()
	// defer s.Close()

	// if !e.ID.Valid() {
	// 	e.ID = bson.NewObjectId()
	// }

	// if !e.Location.ID.Valid() {
	// 	e.Location.ID = bson.NewObjectId()
	// }

	// return []byte(e.ID), s.DB(config.DBName).C(EVENTS).Insert(e)
	return nil, nil
}

func (m *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	// s := m.getFreshSession()
	// defer s.Close()

	// e := persistence.Event{}

	// err := s.DB(config.DBName).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return persistence.Event{}, nil
}

func (m *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	// s := m.getFreshSession()
	// defer s.Close()
	// e := persistence.Event{}
	// err := s.DB(config.DBName).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	// return e, err
	// return events, err
	return persistence.Event{}, nil
}

func (m *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	// s := m.getFreshSession()
	// defer s.Close()
	// events := []persistence.Event{}
	// err := s.DB(config.DBName).C(EVENTS).Find(nil).All(&events)
	// return events, err
	return nil, nil
}
