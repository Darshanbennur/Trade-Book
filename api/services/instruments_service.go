package services

import "github.com/Darshanbennur/Trade-Book/api/models"

type InstruentServices interface {
	// Create a new Instrument : 
	CreateInstrument(*models.Instrument) error

	// Get a given Instrument : 
	Get_SingleInstrument(*string) (*models.Instrument ,error)

	// Get all Instrument : 
	GetAllInstrument() ([] *models.Instrument, error)

	// Update a given Instrument : 
	Update_SingleInstrument(*string) error

	// Delete a given Instrument :
	Delete_SingleInstrument(*string) error
}