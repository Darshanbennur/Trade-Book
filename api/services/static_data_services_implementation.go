package services

import (
	"context"
	"errors"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewCompanyService(db *gorm.DB, ctx context.Context) CompanyServices {
	return &CompanyServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (CS *CompanyServiceImpl) CreateCompany(company *models.Companies) error {
	newID := uuid.New()
	company.ID = newID
	err := CS.db.Create(company).Error
	if err != nil {
		return err
	}
	return nil
}

func (CS *CompanyServiceImpl) Get_SingleCompanyDetails(id *string) (*models.Companies, error) {
	var company models.Companies
	results := CS.db.Where("_id = ?", *id).First(&company)
	if results.Error != nil {
		return nil, results.Error
	}
	return &company, nil
}

func (CS *CompanyServiceImpl) GetAllCompanyDetails() ([]*models.Companies, error) {
	var company []*models.Companies
	result := CS.db.Find(&company)
	if result.Error != nil {
		return nil, result.Error
	}
	return company, nil
}

func (CS *CompanyServiceImpl) Update_SingleCompany(id *string, updatedCompany *models.Companies) error {
	if updatedCompany == nil {
		return errors.New("updated company details cannot be nil")
	}
	var company models.Companies
	result := CS.db.Where("_id = ?", *id).First(&company)
	if result.Error != nil {
		return result.Error
	}
	err := CS.db.Model(&company).Updates(updatedCompany).Error
	if err != nil {
		return err
	}
	return nil
}

func (CS *CompanyServiceImpl) Delete_SingleCompany(id *string) error {
	result := CS.db.Where("_id = ?", *id).Delete(&models.Companies{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
