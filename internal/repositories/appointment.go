// internal/repositories/appointment.go
package repositories

import (
	"gorm.io/gorm"
	"hospital-app/internal/models"
	"time"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *AppointmentRepository) FindAll() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Patient").Preload("Patient.User").
		Preload("Doctor").Preload("Doctor.User").
		Preload("Department").Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) FindByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.db.Preload("Patient").Preload("Patient.User").
		Preload("Doctor").Preload("Doctor.User").
		Preload("Department").First(&appointment, id).Error
	return &appointment, err
}

func (r *AppointmentRepository) Update(appointment *models.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *AppointmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Appointment{}, id).Error
}

func (r *AppointmentRepository) FindByPatient(patientID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Patient").Preload("Patient.User").
		Preload("Doctor").Preload("Doctor.User").
		Preload("Department").
		Where("patient_id = ?", patientID).Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) FindByDoctor(doctorID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Patient").Preload("Patient.User").
		Preload("Doctor").Preload("Doctor.User").
		Preload("Department").
		Where("doctor_id = ?", doctorID).Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) FindByDateRange(start, end time.Time) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Patient").Preload("Patient.User").
		Preload("Doctor").Preload("Doctor.User").
		Preload("Department").
		Where("appointment_date BETWEEN ? AND ?", start, end).Find(&appointments).Error
	return appointments, err
}
