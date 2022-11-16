package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Precription struct {
	ID          primitive.ObjectID `bson:"_id"`
	DoctorID    string             `json:"doctorid" validate:"required"`
	DateAndTime time.Time          `json:"dateandtime" validate:"required"`
	Precription string             `json:"precription" validate:"required"`
}
