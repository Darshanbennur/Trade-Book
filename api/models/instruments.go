package models

import (
	"github.com/google/uuid"
)

type Instrument struct {
	ID             uuid.UUID `json:"_id"`
	InstrumentName string    `json:"instrument_name"`
	Explanation    string    `json:"explanation"`
	VisibilityType string    `json:"visibility_type"`
	SuccessRate    float64   `json:"success_rate"`
}
