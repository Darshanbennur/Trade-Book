package models

import "github.com/google/uuid"

type Companies struct {
	ID            uuid.UUID `gorm:"column:_id;primaryKey"`
	CompanySymbol string    `json:"company_symbol"`
	CompanyName   string    `json:"company_name"`
}
