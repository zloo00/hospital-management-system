package routes

import (
	"github.com/aluazholdykan/hospital-management-system/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	patient := r.Group("/patients")
	{
		patient.GET("/", controllers.GetPatients)
		patient.GET("/:id", controllers.GetPatientByID)
		patient.POST("/", controllers.CreatePatient)
		patient.PUT("/:id", controllers.UpdatePatient)
		patient.DELETE("/:id", controllers.DeletePatient)
	}
	
	doctors := r.Group("/doctors")
	{
		doctors.GET("/", controllers.GetDoctors)
		doctors.GET("/:id", controllers.GetDoctorByID)
		doctors.POST("/", controllers.CreateDoctor)
		doctors.PUT("/:id", controllers.UpdateDoctor)
		doctors.DELETE("/:id", controllers.DeleteDoctor)
	}

}
