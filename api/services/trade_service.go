package services

import "github.com/Darshanbennur/Trade-Book/api/models"

type TradeServices interface {
	// Create a new trade :
	CreateTrade(*models.Trades) error

	// Get a given trade :
	Get_SingleTrade(*string) (*models.Trades, error)

	// Get all the trades :
	GetAllTrade() ([]*models.Trades, error)

	// Update a given trade :
	Update_SingleTradeContent(*models.Trades) error

	// Delete a given trade : 
	Delete_SingleTrade(*string) error
}
