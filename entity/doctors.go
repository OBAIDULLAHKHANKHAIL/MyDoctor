package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct {
	ID                       primitive.ObjectID `bson:"_id"`
	Name                     string             `json:"name" validate:"required,min=3,max=100"`
	Password                 string             `json:"password" validate:"required,min=6"`
	UserName                 string             `json:"username" validate:"required"`
	Address                  []string           `json:"address"`
	PermamentAddress         []string           `json:"permanentaddress"`
	CNIC                     string             `json:"cnic" validate:"required"`
	PassportNo               string             `json:"passportno"`
	PatientID                []Patient          `json:"patientid"`
	Specility                []string           `json:"specility" validate:"required"`
	Gender                   string             `json:"gender" validate:"required"`
	DOB                      time.Time          `json:"dob" validate:"required"`
	MaritalStatus            string             `jaon:"maritalstatus" validate:"required"`
	Experience               []string           `json:"experience" validate:"required"`
	Qualification            []string           `json:"qualification" validate:"required"`
	RegisterationDateAndTime time.Time          `json:"registrationdateandtime" validate:"required"`
	ContactNo                []string           `json:"contactno" validate:"required,min=9,max=12"`
	PersonalPhoneNo          []string           `json:"personalphoneno" validate:"required,min=9,max=12"`
	Timings                  []time.Time        `json:"timing"`
	OnlineStatus             bool               `json:"onlinestatus"`
	Fees                     float32            `json:"fees" validate:"required"`
}
