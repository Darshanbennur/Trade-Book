package services

import "github.com/Darshanbennur/Trade-Book/api/models"

type JournalTradeServices interface {
	// Create a new journal/note :
	CreateJournal(*models.JournalTrades) error

	// Get a journal :
	Get_SingleJournal(*string) (*models.JournalTrades, error)

	// Get all journals :
	GetAllJournal() ([]*models.JournalTrades, error)

	// Update a given journal :
	Update_SingleJournalContent(*string) error

	// Delete a given journal :
	Delete_SingleJournal(*string) error
}
