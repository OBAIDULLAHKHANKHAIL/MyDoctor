package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID                       primitive.ObjectID `bson:"_id"`
	Name                     string             `json:"name" validate:"required,min=3,max=100"`
	Address                  []string           `json:"address"`
	PermamentAddress         []string           `json:"permanentaddress" validate:"required"`
	UserName                 string             `json:"username" validate:"required"`
	Password                 string             `json:"password" validate:"required,min=6"`
	CNIC                     string             `json:"cnic" validate:"required"`
	PassportNo               string             `json:"passportno"`
	RegisterationDateAndTime time.Time          `json:"registrationdateandtime" validate:"required"`
	ContactNo                []string           `json:"contactno" validate:"required"`
	PersonalPhoneNo          []string           `json:"personalphoneno"`
	DOB                      time.Time          `json:"dob" validate:"required"`
	Gender                   string             `json:"gender" validate:"required"`
	MaritalStatus            string             `jaon:"maritalstatus"`
	DoctorID                 []string           `json:"doctorid"`
	PreciptionID             []string           `json:"precriptionid"`
	ReportsID                []string           `json:"reportsid"`
	BillID                   []string           `json:"bill"`
}
