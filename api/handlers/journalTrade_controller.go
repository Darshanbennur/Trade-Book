package handlers

import (
	"net/http"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/Darshanbennur/Trade-Book/api/services"
	"github.com/gin-gonic/gin"
)

type JournalTradeController struct {
	JournalTradeService services.JournalTradeServices
}

func NewJournalTrade(journalTradeService services.JournalTradeServices) JournalTradeController {
	return JournalTradeController{
		JournalTradeService: journalTradeService,
	}
}

func (JC *JournalTradeController) CreateJournalTrade(ctx *gin.Context) {
	var journal models.JournalTrades
	if err := ctx.ShouldBindJSON(&journal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := JC.JournalTradeService.CreateJournal(&journal)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (JC *JournalTradeController) Get_SingleJournal(ctx *gin.Context) {
	journalTradeID := ctx.Param("id")
	trade, err := JC.JournalTradeService.Get_SingleJournal(&journalTradeID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"instrument": trade})
}

func (JC *JournalTradeController) GetAllJournal(ctx *gin.Context) {
	journals, err := JC.JournalTradeService.GetAllJournal()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"journals": journals})
}

func (JC *JournalTradeController) Update_SingleJournalContent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Updation Successful"})
}

func (JC *JournalTradeController) Delete_SingleJournal(ctx *gin.Context) {
	journalTradeID := ctx.Param("id")
	err := JC.JournalTradeService.Delete_SingleJournal(&journalTradeID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deletion Successful"})
}

func (JC *JournalTradeController) RegisterJournalTradeRoutes(rg *gin.RouterGroup) {
	tradeRoute := rg.Group("/journal")
	tradeRoute.POST("/create", JC.CreateJournalTrade)
	tradeRoute.GET("/get/:id", JC.Get_SingleJournal)
	tradeRoute.GET("/getAll", JC.GetAllJournal)
	tradeRoute.PATCH("/update", JC.Update_SingleJournalContent)
	tradeRoute.DELETE("/delete/:id", JC.Delete_SingleJournal)
}
