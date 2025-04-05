package controllers

import (
	"github.com/aluazholdykan/hospital-management-system/config"
	"github.com/aluazholdykan/hospital-management-system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB = config.DB

func GetDoctors(c *gin.Context) {
	var doctors []models.Doctor
	// Check for any error during the query
	if err := db.Find(&doctors).Error; err != nil {
		log.Println("Error fetching doctors:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching doctors"})
		return
	}
	c.JSON(http.StatusOK, doctors)
}

func GetDoctorByID(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor
	if err := db.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}
	c.JSON(http.StatusOK, doctor)
}

func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	if err := db.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating doctor"})
		return
	}
	c.JSON(http.StatusCreated, doctor)
}

func UpdateDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor
	if err := db.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating doctor"})
		return
	}
	c.JSON(http.StatusOK, doctor)
}

func DeleteDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor
	if err := db.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}
	if err := db.Delete(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting doctor"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Доктор удалён"})
}
