package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name" validate:"required,min=3,max=100"`
	UserName string             `json:"userName" validate:"required,min=3,max=100"`
	Password string             `json:"Password" validate:"required,min=6"`
}
