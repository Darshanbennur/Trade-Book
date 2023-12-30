package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TradeServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewTradeService(db *gorm.DB, ctx context.Context) TradeServices {
	return &TradeServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (TS *TradeServiceImpl) CreateTrade(trade *models.Trades) error {
	newID := uuid.New()
	currentDate := time.Now()
	trade.ID = newID
	trade.EnterDate = currentDate
	fmt.Printf("Trade : %+v", trade)
	err := TS.db.Create(trade).Error
	if err != nil {
		return err
	}
	return nil
}

func (TS *TradeServiceImpl) Get_SingleTrade(ID *string) (*models.Trades, error) {
	var trade models.Trades
	result := TS.db.Where("_id = ?", *ID).First(&trade)
	if result.Error != nil {
		return nil, result.Error
	}
	return &trade, nil
}

func (TS *TradeServiceImpl) GetAllTrade() ([]*models.Trades, error) {
	var trades []*models.Trades
	result := TS.db.Find(&trades)
	if result.Error != nil {
		return nil, result.Error
	}
	return trades, nil
}

func (TS *TradeServiceImpl) Update_SingleTradeContent(trade *models.Trades) error {
	result := TS.db.Save(trade)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (TS *TradeServiceImpl) Delete_SingleTrade(ID *string) error {
	result := TS.db.Where("_id = ?", *ID).Delete(&models.Trades{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
