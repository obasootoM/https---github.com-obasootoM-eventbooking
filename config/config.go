package config

import (
	"encoding/json"
	dblayer "eventsbook/dbLayer"
	"log"
	"os"
)

const (
	DATACONNECTION    = dblayer.MongoDb
	DEFAULTCONNECTION = "mongodb://127.0.0.1"
	RESTCONNECTION    = "localhost:8585"
)

type Config struct {
	DatabaseConnection dblayer.DataType `json:"databaseonnection"`
	DefaultConnection  string           `json:"defaultconnection"`
	RESTCONNECTION     string            `json:"restconnection"`
}

func NewConfig(connection string) (*Config, error) {
	config := &Config{
      DATACONNECTION,
	  DEFAULTCONNECTION,
	  RESTCONNECTION,
	}
	fileName, err := os.Open(connection)
	if err != nil {
		log.Fatal("cannot open connection")
		return config,err
	}
	defer fileName.Close()
	decode := json.NewDecoder(fileName)
	err = decode.Decode(&config)
	return config, err
	
}

