package main

import (
	"github.com/eminoz/rabbitmq/broker"
	"github.com/eminoz/rabbitmq/pkg/config"
	"github.com/eminoz/rabbitmq/pkg/database"
)

func main() {
	config.SetupConfig()
	database.SetDatabase()

	broker.Consumers()

}
