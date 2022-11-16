package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Report struct {
	ID          primitive.ObjectID `bson:"_id"`
	DoctorID    string             `json:"doctorid" validate:"required"`
	DateAndTime time.Time          `json:"updated_at" validate:"required"`
	Laboratory  string             `json:"laboratory" validate:"required"`
	Report      string             `json:"report" validate:"required"`
}
