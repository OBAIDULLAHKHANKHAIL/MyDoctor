package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bill struct {
	ID          primitive.ObjectID `bson:"_id"`
	Amount      float32            `json:"amount" validate:"required"`
	Status      string             `json:"status" validate:"required"`
	DoctorID    string             `json:"doctorid" validate:"required"`
	DateAndTime time.Time          `json:"dateandtime" validate:"required"`
}
