package database

import (
	"eventsbook/persistence"
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB    = "myevent"
	USER  = "user"
	EVENT = "event"
)

type Mongo struct {
	mongo *mgo.Session
}

func NewMongo(connection string) (*Mongo, error) {
	m, err := mgo.Dial(connection)
	if err != nil {
		fmt.Print("cannot connect to database")
		os.Exit(1)
	}
	return &Mongo{
		mongo: m,
	}, err
}

func (m *Mongo) getFreshMongo() *mgo.Session {
	return m.mongo.Copy()
}

func (m *Mongo) AddEvent(e persistence.Event) ([]byte, error) {
	s := m.getFreshMongo()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}
	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}
	return []byte(e.ID), s.DB(DB).C(EVENT).Insert()
}
func (m *Mongo) FindEvent(id []byte) (error, persistence.Event) {

}
func (m *Mongo) FindAllEventAvailable() (error, []persistence.Event) {

}
func (m *Mongo) FindEventByName(name string) (error, persistence.Event) {

}
