package controllers

import (
	"hospital-app/internal/models"
	"hospital-app/internal/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppointmentController struct {
	appointmentRepo *repositories.AppointmentRepository
	patientRepo     *repositories.PatientRepository
	doctorRepo      *repositories.DoctorRepository
	departmentRepo  *repositories.DepartmentRepository
}

func NewAppointmentController(db *gorm.DB) *AppointmentController {
	return &AppointmentController{
		appointmentRepo: repositories.NewAppointmentRepository(db),
		patientRepo:     repositories.NewPatientRepository(db),
		doctorRepo:      repositories.NewDoctorRepository(db),
		departmentRepo:  repositories.NewDepartmentRepository(db),
	}
}

type CreateAppointmentRequest struct {
	PatientID       uint      `json:"patientId" binding:"required"`
	DoctorID        uint      `json:"doctorId" binding:"required"`
	DepartmentID    uint      `json:"departmentId" binding:"required"`
	AppointmentDate time.Time `json:"appointmentDate" binding:"required"`
}

func (ctrl *AppointmentController) CreateAppointment(c *gin.Context) {
	var req CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if patient exists
	_, err := ctrl.patientRepo.FindByID(req.PatientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// Check if doctor exists
	_, err = ctrl.doctorRepo.FindByID(req.DoctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor not found"})
		return
	}

	// Check if department exists
	_, err = ctrl.departmentRepo.FindByID(req.DepartmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
		return
	}

	appointment := models.Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		DepartmentID:    req.DepartmentID,
		AppointmentDate: req.AppointmentDate,
		Status:          "scheduled",
	}

	if err := ctrl.appointmentRepo.Create(&appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func (ctrl *AppointmentController) GetAppointments(c *gin.Context) {
	appointments, err := ctrl.appointmentRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ctrl *AppointmentController) GetAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	appointment, err := ctrl.appointmentRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ctrl *AppointmentController) UpdateAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	appointment, err := ctrl.appointmentRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.appointmentRepo.Update(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ctrl *AppointmentController) DeleteAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	if err := ctrl.appointmentRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "appointment deleted successfully"})
}

func (ctrl *AppointmentController) GetPatientAppointments(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patientId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID format"})
		return
	}

	appointments, err := ctrl.appointmentRepo.FindByPatient(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (ctrl *AppointmentController) GetDoctorAppointments(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctorId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID format"})
		return
	}

	appointments, err := ctrl.appointmentRepo.FindByDoctor(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

type UpdateAppointmentStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

func (ctrl *AppointmentController) UpdateAppointmentStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var req UpdateAppointmentStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment, err := ctrl.appointmentRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	appointment.Status = req.Status

	if err := ctrl.appointmentRepo.Update(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

type AddDiagnosisRequest struct {
	Diagnosis    string `json:"diagnosis" binding:"required"`
	Prescription string `json:"prescription"`
	Notes        string `json:"notes"`
}

func (ctrl *AppointmentController) AddDiagnosis(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var req AddDiagnosisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment, err := ctrl.appointmentRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	appointment.Diagnosis = req.Diagnosis
	appointment.Prescription = req.Prescription
	appointment.Notes = req.Notes
	appointment.Status = "completed"

	if err := ctrl.appointmentRepo.Update(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}
