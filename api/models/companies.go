package api

import "github.com/google/uuid"

type Companies struct {
	ID            uuid.UUID `json:"_id"`
	CompanySymbol string    `json:"company_symbol"`
	CompanyName   string    `json:"company_name"`
}
