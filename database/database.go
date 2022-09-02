package database

import (
	"eventsbook/persistence"
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)


type Mongo struct {
	mongo *mgo.Session
}

func NewMongo(connection string) ( *Mongo,error) {
  m, err := mgo.Dial(connection)
  if err != nil {
	fmt.Print("cannot connect to database")
	os.Exit(1)
  }
  return &Mongo{
	mongo: m,
  },err
}

func (m *Mongo) getFreshMongo() *mgo.Session {
   return m.mongo.Copy()
}

func (m *Mongo) AddEvent(e persistence.Event) ([]byte, error) {
   
}
func (m *Mongo) FindEvent(id []byte)( error, persistence.Event) {

}
func (m *Mongo) FindAllEventAvailable() (error, []persistence.Event){

}
func (m *Mongo) FindEventByName(name string) (error, persistence.Event) {

}