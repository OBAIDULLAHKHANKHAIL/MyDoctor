package main

import (
	"github.com/obaidullahkhankhail/MyDoctor/controller"
	"github.com/obaidullahkhankhail/MyDoctor/service"

	"github.com/gin-gonic/gin"
)

var (
	PatientService    service.PatientService       = service.NewPatient()
	PatientController controller.PatientController = controller.NewPatient(PatientService)

	DoctorService    service.DoctorService       = service.NewDoctor()
	DoctorController controller.DoctorController = controller.NewDoctor(DoctorService)
)

func main() {
	server := gin.Default()
	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message":"hi this is my code",
	// 	})
	// })

	server.GET("/patients", func(ctx *gin.Context) {
		ctx.JSON(200, PatientController.FindAllPatients())
	})

	server.POST("/patients", func(ctx *gin.Context) {
		ctx.JSON(200, PatientController.RegisterPatient(ctx))
	})

	server.GET("/doctors", func(ctx *gin.Context) {
		ctx.JSON(200, PatientController.FindAllPatients())
	})

	server.POST("/doctors", func(ctx *gin.Context) {
		ctx.JSON(200, PatientController.RegisterPatient(ctx))
	})

	server.Run(":8081")
}
