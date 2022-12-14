package main

import (
	configuration "eventsbook/config"
	dblayer "eventsbook/dbLayer"
	"eventsbook/lib/kafka"
	service "eventsbook/service"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

func main() {
	path := flag.String("conf", `.\configuration\config.json`, "set config path to json")
	flag.Parse()
	config, _ := configuration.NewConfig(*path)
	fmt.Println("connecting to database")
	//all connection
	//kafka configuration
	brokerList := os.Getenv("KAFKA_BROKER")
	if brokerList == "" {
		brokerList = "localhost:9092"
	}
	brokers := strings.Split(brokerList, ",")
	configurate := sarama.NewConfig()
	client, err := sarama.NewClient(brokers, configurate)
	if err != nil {
		panic(err)
	}
	produce, err := kafka.NewProducer(client)
	if err != nil {
		panic(err)
	}

	dbHandler, _ := dblayer.NewPersistent(config.DataType, config.DefaultConnection)
	httpTLS, httpAPI := service.Server(config.RestApi, config.TlsApi, dbHandler, produce)
	select {
	case err := <-httpAPI:
		if err != nil {
			fmt.Println("cannot load http")
		}
	case err := <-httpTLS:
		if err != nil {
			fmt.Println("cannot load https")
		}
	}
}
