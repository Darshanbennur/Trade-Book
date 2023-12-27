package handlers

import (
	"net/http"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/Darshanbennur/Trade-Book/api/services"
	"github.com/gin-gonic/gin"
)

type InstrumentController struct {
	InstrumentService services.InstrumentServices
}

func NewInstrument(instrumentServices services.InstrumentServices) InstrumentController {
	return InstrumentController{
		InstrumentService: instrumentServices,
	}
}

func (is *InstrumentController) CreateInstrument(ctx *gin.Context) {
	var instrument models.Instrument
	if err := ctx.ShouldBindJSON(&instrument); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := is.InstrumentService.CreateInstrument(&instrument)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (is *InstrumentController) Get_SingleInstrument(ctx *gin.Context) {
	instrumentID := ctx.Param("name")
	instrument, err := is.InstrumentService.Get_SingleInstrument(&instrumentID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"instrument": instrument})
}

func (is *InstrumentController) GetAllInstrument(ctx *gin.Context) {

}

func (is *InstrumentController) Update_SingleInstrument(ctx *gin.Context) {

}

func (is *InstrumentController) Delete_SingleInstrument(ctx *gin.Context) {

}

func (is *InstrumentController) RegisterInstrumentRoutes(rg *gin.RouterGroup) {
	instrumentRoute := rg.Group("/instrument")
	instrumentRoute.POST("/create", is.CreateInstrument)
	instrumentRoute.GET("/get/:name", is.Get_SingleInstrument)
	instrumentRoute.GET("/getAll", is.GetAllInstrument)
	instrumentRoute.PATCH("/update", is.Update_SingleInstrument)
	instrumentRoute.DELETE("/delete/:name", is.Delete_SingleInstrument)
}
