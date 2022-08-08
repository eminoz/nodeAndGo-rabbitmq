package service

import (
	"context"
	"fmt"
	"github.com/eminoz/rabbitmq/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUser interface {
	CreateUser(ctx context.Context, user *model.User) *mongo.InsertOneResult
}

func (c *UserCollection) CreateUser(ctx context.Context, user *model.User) *mongo.InsertOneResult {
	fmt.Println(user)
	one, _ := c.Collection.InsertOne(ctx, &user)
	return one
}
