package controllers

import (
	"hospital-app/internal/models"
	"hospital-app/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DoctorController struct {
	doctorRepo     *repositories.DoctorRepository
	userRepo       *repositories.UserRepository
	departmentRepo *repositories.DepartmentRepository
}

func NewDoctorController(db *gorm.DB) *DoctorController {
	return &DoctorController{
		doctorRepo:     repositories.NewDoctorRepository(db),
		userRepo:       repositories.NewUserRepository(db),
		departmentRepo: repositories.NewDepartmentRepository(db),
	}
}

func (ctrl *DoctorController) CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	_, err := ctrl.userRepo.FindByID(doctor.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// Check if department exists
	if doctor.DepartmentID != 0 {
		_, err := ctrl.departmentRepo.FindByID(doctor.DepartmentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
			return
		}
	}

	if err := ctrl.doctorRepo.Create(&doctor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

func (ctrl *DoctorController) GetDoctors(c *gin.Context) {
	doctors, err := ctrl.doctorRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

func (ctrl *DoctorController) GetDoctor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID"})
		return
	}

	doctor, err := ctrl.doctorRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "doctor not found"})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func (ctrl *DoctorController) UpdateDoctor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID"})
		return
	}

	doctor, err := ctrl.doctorRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "doctor not found"})
		return
	}

	if err := c.ShouldBindJSON(doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if department exists
	if doctor.DepartmentID != 0 {
		_, err := ctrl.departmentRepo.FindByID(doctor.DepartmentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "department not found"})
			return
		}
	}

	if err := ctrl.doctorRepo.Update(doctor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, doctor)
}

func (ctrl *DoctorController) DeleteDoctor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID"})
		return
	}

	if err := ctrl.doctorRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "doctor deleted successfully"})
}

func (ctrl *DoctorController) GetDoctorsByDepartment(c *gin.Context) {
	departmentIDStr := c.Param("departmentId")
	departmentID, err := strconv.ParseUint(departmentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid department ID"})
		return
	}

	doctors, err := ctrl.doctorRepo.FindByDepartment(uint(departmentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, doctors)
}
