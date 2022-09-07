package configuration

import (
	"encoding/json"
	dblayer "eventsbook/dbLayer"
	"fmt"
	"os"
)

var (
	DATATYPES         = dblayer.MONGODB
	DEFAULTCONNECTION = "mongodb://127.0.0.1"
	RESTAPI           = "localhost:8181"
	TLSAPI            = "localhost:8282"
)

type Config struct {
	DataType          dblayer.DATATYPE `json:"datatype"`
	DefaultConnection string           `json:"defaultconnection"`
	RestApi           string           `json:"rest_api"`
	TlsApi            string           `json:"tls_api"`
}

func NewConfig(connection string) (*Config, error) {
	config := &Config{
		DATATYPES,
		DEFAULTCONNECTION,
		RESTAPI,
		TLSAPI,
	}
	fileName, err := os.Open(connection)
	if err != nil {
		fmt.Println("cannot open connection, can still continue")
		return config, err
	}
	defer fileName.Close()
	decode := json.NewDecoder(fileName)
	err = decode.Decode(&config)
	if err != nil {
		fmt.Println("cannot decode file")
	}
	return config, err

}
