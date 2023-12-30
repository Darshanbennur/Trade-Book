package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Trades struct {
	ID         uuid.UUID       `gorm:"column:_id;primaryKey"`
	CompanyID  uuid.UUID       `json:"company_id"` // will refer to the _id(UUID) of the Company Traded
	EnterDate  time.Time       `json:"enter_date"`
	Pnl        float64         `json:"pnl"`
	Instrument json.RawMessage `json:"instrument"` // will refer to the _id(UUID) of the instrument Used
	Side       string          `json:"side"`
	Volume     int64           `json:"volume"`
	NetRoi     float64         `json:"net_roi"`
	Notes      json.RawMessage `json:"notes"` //references the _id of the "JournalTrades" associated with this trade.
}
