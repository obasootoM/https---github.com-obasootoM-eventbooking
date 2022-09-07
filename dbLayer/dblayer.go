package dblayer

import (
	"eventsbook/database"
	"eventsbook/persistence"
)

type DATATYPE string
const(
   MONGODB DATATYPE = "mongodb"
)

func NewPersistent(option DATATYPE, connection string) (persistence.Database, error) {
    switch option {
	case MONGODB:
		return database.NewMongo(connection)
	}
	return nil, nil
}

