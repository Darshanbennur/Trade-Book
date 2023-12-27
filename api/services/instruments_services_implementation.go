package services

import (
	"context"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InstrumentServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewInstrumentService(db *gorm.DB, ctx context.Context) InstrumentServices {
	return &InstrumentServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (IS *InstrumentServiceImpl) CreateInstrument(instrument *models.Instrument) error {
	newID := uuid.New()
	instrument.ID = newID
	err := IS.db.Create(instrument).Error
	if err != nil {
		return err
	}
	return nil
}

func (IS *InstrumentServiceImpl) Get_SingleInstrument(name *string) (*models.Instrument, error) {
	var instrument models.Instrument
	result := IS.db.Where("instrument_name = ?", *name).First(&instrument)
	if result.Error != nil {
		return nil, result.Error
	}
	return &instrument, nil
}

func (IS *InstrumentServiceImpl) GetAllInstrument() ([]*models.Instrument, error) {
	panic("unimplemented")
}

func (IS *InstrumentServiceImpl) Update_SingleInstrument(*models.Instrument) error {
	panic("unimplemented")
}

func (IS *InstrumentServiceImpl) Delete_SingleInstrument(*string) error {
	panic("unimplemented")
}