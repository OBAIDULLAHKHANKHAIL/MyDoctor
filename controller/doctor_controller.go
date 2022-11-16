package controller

import (
	"github.com/obaidullahkhankhail/MyDoctor/entity"
	"github.com/obaidullahkhankhail/MyDoctor/service"

	"github.com/gin-gonic/gin"
)

type DoctorController interface {
	FindAllDoctors() []entity.Doctor
	RegisterDoctor(ctx *gin.Context) entity.Doctor
}

type controller struct {
	service service.DoctorService
}

func NewDoctor(service service.DoctorService) DoctorController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAllDoctors() []entity.Doctor {
	return c.service.FindAllDoctors()
}

func (c *controller) RegisterDoctor(ctx *gin.Context) entity.Doctor {
	var doctor entity.Doctor
	ctx.BindJSON(&doctor)
	c.service.RegisterDoctor(doctor)
	return doctor
}
