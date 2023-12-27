package models

import (
	"time"

	"github.com/google/uuid"
)

type Trades struct {
	ID         uuid.UUID   `gorm:"column:_id;primaryKey"`
	CompanyID  uuid.UUID   `json:"company_id"` // will refer to the _id(UUID) of the Company Traded
	Date       time.Time   `json:"enter_date"`
	PnL        float64     `json:"pnl"`
	Instrument []uuid.UUID `json:"instrument"` // will refer to the _id(UUID) of the instrument Used
	Side       string      `json:"side"`
	Volume     int64       `json:"volume"`
	NetROI     float64     `json:"net_roi"`
	Notes      []uuid.UUID `json:"notes"` //references the _id of the "JournalTrades" associated with this trade.
}
