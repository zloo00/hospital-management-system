package repositories

import (
	"gorm.io/gorm"
	"hospital-app/internal/models"
)

type DoctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) Create(doctor *models.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *DoctorRepository) FindAll() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := r.db.Preload("User").Preload("Department").Find(&doctors).Error
	return doctors, err
}

func (r *DoctorRepository) FindByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.Preload("User").Preload("Department").First(&doctor, id).Error
	return &doctor, err
}

func (r *DoctorRepository) Update(doctor *models.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *DoctorRepository) Delete(id uint) error {
	return r.db.Delete(&models.Doctor{}, id).Error
}

func (r *DoctorRepository) FindByDepartment(departmentID uint) ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := r.db.Preload("User").Where("department_id = ?", departmentID).Find(&doctors).Error
	return doctors, err
}
