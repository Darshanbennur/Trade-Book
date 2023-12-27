package models

import (
	"github.com/google/uuid"
)

type Instrument struct {
	ID             uuid.UUID `gorm:"column:_id;primaryKey"`
	InstrumentName string    `json:"instrument_name"`
	Explanation    string    `json:"explanation"`
	VisibilityType string    `json:"visibility_type"`
	SuccessRate    float64   `json:"success_rate"`
}
