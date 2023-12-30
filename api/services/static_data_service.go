package services

import "github.com/Darshanbennur/Trade-Book/api/models"

/*
Will contain services of static components like :
	- CRUD operations on Company Table
	- Other Operations as well (NOTE : enlist which are used here)
*/

type CompanyServices interface {
	// Create a new Company
	CreateCompany(*models.Companies) error

	// Get a company :
	Get_SingleCompanyDetails(*string) (*models.Companies, error)

	// Get all company details :
	GetAllCompanyDetails() ([]*models.Companies, error)

	// Update a given company :
	Update_SingleCompany(*string, *models.Companies) error

	// Delete a given company :
	Delete_SingleCompany(*string) error
}
