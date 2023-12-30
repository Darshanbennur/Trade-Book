package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JournalTradeServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewJournalTradeService(db *gorm.DB, ctx context.Context) JournalTradeServices {
	return &JournalTradeServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (JS *JournalTradeServiceImpl) CreateJournal(journal *models.JournalTrades) error {
	newID := uuid.New()
	currentDate := time.Now()
	journal.ID = newID
	journal.EnterDate = currentDate
	err := JS.db.Create(journal).Error
	if err != nil {
		return err
	}
	return nil
}

func (JS *JournalTradeServiceImpl) Get_SingleJournal(id *string) (*models.JournalTrades, error) {
	var journal models.JournalTrades
	result := JS.db.Where("_id = ?", *id).First(&journal)
	if result.Error != nil {
		return nil, result.Error
	}
	return &journal, nil
}

func (JS *JournalTradeServiceImpl) GetAllJournal() ([]*models.JournalTrades, error) {
	var journals []*models.JournalTrades
	result := JS.db.Find(&journals)
	if result.Error != nil {
		return nil, result.Error
	}
	return journals, nil
}

func (JS *JournalTradeServiceImpl) Update_SingleJournalContent(id *string) error {
	fmt.Println("Updating journal for ID:", id)
	return nil
}

func (JS *JournalTradeServiceImpl) Delete_SingleJournal(id *string) error {
	result := JS.db.Where("_id = ?", *id).Delete(&models.JournalTrades{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
