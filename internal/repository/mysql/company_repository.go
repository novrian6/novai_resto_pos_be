package mysql

import (
	"resto-pos-backend/internal/domain/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.DB.Create(company).Error
}

func (r *CompanyRepository) GetByID(id int) (*models.Company, error) {
	var company models.Company
	err := r.DB.First(&company, id).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Update(company *models.Company) error {
	return r.DB.Save(company).Error
}

func (r *CompanyRepository) Delete(id int) error {
	return r.DB.Delete(&models.Company{}, id).Error
}
