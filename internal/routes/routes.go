package routes

import (
	"hospital-app/internal/config"
	"hospital-app/internal/controllers"
	"hospital-app/internal/middleware"
	"hospital-app/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	jwtUtils := utils.NewJWTUtils(cfg.JWTSecret)
	authMiddleware := middleware.AuthMiddleware(jwtUtils)

	// Auth routes
	authController := controllers.NewAuthController(db, cfg.JWTSecret)
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}

	// Admin routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(authMiddleware)
	adminGroup.Use(middleware.RoleMiddleware("admin"))
	{
		patientController := controllers.NewPatientController(db)
		adminGroup.GET("/patients", patientController.GetPatients)
		adminGroup.GET("/patients/:id", patientController.GetPatient)
		adminGroup.POST("/patients", patientController.CreatePatient)
		adminGroup.PUT("/patients/:id", patientController.UpdatePatient)
		adminGroup.DELETE("/patients/:id", patientController.DeletePatient)

		doctorController := controllers.NewDoctorController(db)
		adminGroup.GET("/doctors", doctorController.GetDoctors)
		adminGroup.GET("/doctors/:id", doctorController.GetDoctor)
		adminGroup.POST("/doctors", doctorController.CreateDoctor)
		adminGroup.PUT("/doctors/:id", doctorController.UpdateDoctor)
		adminGroup.DELETE("/doctors/:id", doctorController.DeleteDoctor)

		departmentController := controllers.NewDepartmentController(db)
		adminGroup.GET("/departments", departmentController.GetDepartments)
		adminGroup.GET("/departments/:id", departmentController.GetDepartment)
		adminGroup.POST("/departments", departmentController.CreateDepartment)
		adminGroup.PUT("/departments/:id", departmentController.UpdateDepartment)
		adminGroup.DELETE("/departments/:id", departmentController.DeleteDepartment)

		appointmentController := controllers.NewAppointmentController(db)
		adminGroup.GET("/appointments", appointmentController.GetAppointments)
		adminGroup.GET("/appointments/:id", appointmentController.GetAppointment)
		adminGroup.POST("/appointments", appointmentController.CreateAppointment)
		adminGroup.PUT("/appointments/:id", appointmentController.UpdateAppointment)
		adminGroup.DELETE("/appointments/:id", appointmentController.DeleteAppointment)
	}

	// Doctor routes
	doctorGroup := router.Group("/doctor")
	doctorGroup.Use(authMiddleware)
	doctorGroup.Use(middleware.RoleMiddleware("doctor"))
	{
		appointmentController := controllers.NewAppointmentController(db)
		doctorGroup.GET("/appointments", appointmentController.GetDoctorAppointments)
		doctorGroup.PUT("/appointments/:id/status", appointmentController.UpdateAppointmentStatus)
		doctorGroup.PUT("/appointments/:id/diagnosis", appointmentController.AddDiagnosis)

		patientController := controllers.NewPatientController(db)
		doctorGroup.GET("/patients", patientController.GetPatients)
		doctorGroup.GET("/patients/:id", patientController.GetPatient)
	}

	// Patient routes
	patientGroup := router.Group("/patient")
	patientGroup.Use(authMiddleware)
	patientGroup.Use(middleware.RoleMiddleware("patient"))
	{
		appointmentController := controllers.NewAppointmentController(db)
		patientGroup.GET("/appointments", appointmentController.GetPatientAppointments)
		patientGroup.POST("/appointments", appointmentController.CreateAppointment)

		doctorController := controllers.NewDoctorController(db)
		patientGroup.GET("/doctors", doctorController.GetDoctors)
		patientGroup.GET("/departments/:departmentId/doctors", doctorController.GetDoctorsByDepartment)

		departmentController := controllers.NewDepartmentController(db)
		patientGroup.GET("/departments", departmentController.GetDepartments)
	}

	return router
}
