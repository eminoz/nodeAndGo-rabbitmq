package broker

import (
	"github.com/eminoz/rabbitmq/controller"
	"github.com/eminoz/rabbitmq/service"
)

func Consumers() {
	collectionSetting := service.UserCollectionSetting()
	userController := controller.UserController{UserService: collectionSetting}
	userConsumer(userController)
}
