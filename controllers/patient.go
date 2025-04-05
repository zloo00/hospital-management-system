package controllers

import (
	"github.com/aluazholdykan/hospital-management-system/config"
	"github.com/aluazholdykan/hospital-management-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&patient)
	c.JSON(http.StatusCreated, patient)
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	config.DB.Delete(&patient)
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
