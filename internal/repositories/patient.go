package repositories

import (
	"gorm.io/gorm"
	"hospital-app/internal/models"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *PatientRepository) FindAll() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Preload("User").Find(&patients).Error
	return patients, err
}

func (r *PatientRepository) FindByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Preload("User").First(&patient, id).Error
	return &patient, err
}

func (r *PatientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *PatientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}
