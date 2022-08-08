package controller

import (
	"context"
	"fmt"
	"github.com/eminoz/rabbitmq/model"
	"github.com/eminoz/rabbitmq/service"
)

type IUserController interface {
	SaveUser(user *model.User)
}
type UserController struct {
	UserService service.IUser
}

func (u *UserController) SaveUser(user *model.User) {
	var ctx context.Context
	createUser := u.UserService.CreateUser(ctx, user)
	fmt.Println(createUser.InsertedID)
}
