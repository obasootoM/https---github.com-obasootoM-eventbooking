package persistence

import "gopkg.in/mgo.v2/bson"


type Database interface {
	AddEvent(Event) ([]byte, error) 
	FindEvent([]byte) (error, Event)
	FindEventByName(string) (error, Event)
	FindAllEventAvailable() (error, []Event)
}

type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  string
	CloseTime string
	Hal       []Hall
}

type Hall struct {
	Name     string 
	Location string 
	Capacity int    
}