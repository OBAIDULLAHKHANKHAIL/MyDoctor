package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID          primitive.ObjectID `bson:"_id"`
	DoctorID    string             `json:"doctorid" validate:"required"`
	PatientID   string             `json:"patientid" validate:"required"`
	DateAndTime time.Time          `json:"dateandtime" validate:"required"`
}
