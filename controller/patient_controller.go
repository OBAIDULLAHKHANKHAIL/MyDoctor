package controller

import (
	"github.com/obaidullahkhankhail/MyDoctor/entity"
	"github.com/obaidullahkhankhail/MyDoctor/service"

	"github.com/gin-gonic/gin"
)

type PatientController interface {
	FindAllPatients() []entity.Patient
	RegisterPatient(ctx *gin.Context) entity.Patient
}

type patient_controller struct {
	service service.PatientService
}

func NewPatient(service service.PatientService) PatientController {
	return &patient_controller{
		service: service,
	}
}

func (c *patient_controller) FindAllPatients() []entity.Patient {
	return c.service.FindAllPatients()
}

func (c *patient_controller) RegisterPatient(ctx *gin.Context) entity.Patient {
	var patient entity.Patient
	ctx.BindJSON(&patient)
	c.service.RegisterPatient(patient)
	return patient
}
