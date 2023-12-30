package models

import (
	"time"

	"github.com/google/uuid"
)

type JournalTrades struct {
	ID      uuid.UUID `gorm:"column:_id;primaryKey"`
	TradeID uuid.UUID `json:"trade_id"` // References the _id from the "Trades" table
	Date    time.Time `json:"enter_date"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}
