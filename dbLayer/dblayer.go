package dblayer

import (
	"eventsbook/database"
	"eventsbook/persistence"
)

type DataType string
const(
   MongoDb DataType = "mongodb"
)

func NewPesistent(connection string,option DataType) (persistence.Database, error) {
    switch option {
	case MongoDb:
		return database.NewMongo(connection)
	}
	return nil, nil
}

