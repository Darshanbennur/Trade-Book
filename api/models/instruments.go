package api

import (
	"github.com/google/uuid"
)

type Instrument struct {
	ID             uuid.UUID `json:"_id"`
	InstrumentName string    `json:"instrument_name"`
	SuccessRate    float64   `json:"success_rate"`
}
