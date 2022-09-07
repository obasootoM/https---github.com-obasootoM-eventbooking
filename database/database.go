package database

import (
	"eventsbook/persistence"

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
		return nil, err
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
func (m *Mongo) FindEvent(id []byte) (persistence.Event, error) {
	s := m.getFreshMongo().Copy()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENT).FindId(bson.ObjectId(id)).One(e)
	return e, err

}
func (m *Mongo) FindAllEventAvailable() ([]persistence.Event, error) {
	s := m.getFreshMongo().Copy()
	defer s.Close()
	e := []persistence.Event{}
	err := s.DB(DB).C(EVENT).Find(nil).All(&e)
	return e, err
}
func (m *Mongo) FindEventByName(name string) (persistence.Event, error) {
	s := m.getFreshMongo().Copy()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENT).Find(bson.M{"name": name}).One(&e)
	return e, err
}
