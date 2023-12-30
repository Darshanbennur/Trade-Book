package handlers

import (
	"net/http"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/Darshanbennur/Trade-Book/api/services"
	"github.com/gin-gonic/gin"
)

type TradeController struct {
	TradeService services.TradeServices
}

func NewTrade(tradeService services.TradeServices) TradeController {
	return TradeController{
		TradeService: tradeService,
	}
}

func (TC *TradeController) CreateTrade(ctx *gin.Context) {
	var trade models.Trades
	if err := ctx.ShouldBindJSON(&trade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := TC.TradeService.CreateTrade(&trade)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (TC *TradeController) Get_SingleTrade(ctx *gin.Context) {
	tradeID := ctx.Param("name")
	trade, err := TC.TradeService.Get_SingleTrade(&tradeID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"instrument": trade})
}

func (TC *TradeController) GetAllTrade(ctx *gin.Context) {
	trades, err := TC.TradeService.GetAllTrade()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"instrument": trades})
}

func (TC *TradeController) Update_SingleTradeContent(ctx *gin.Context) {
	var trade models.Trades
	if err := ctx.ShouldBindJSON(&trade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	erro := TC.TradeService.Update_SingleTradeContent(&trade)
	if erro != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": erro.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updation Successful"})
}

func (TC *TradeController) Delete_SingleTrade(ctx *gin.Context) {
	tradeID := ctx.Param("_id")
	err := TC.TradeService.Delete_SingleTrade(&tradeID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deletion Successful"})
}

func (TC *TradeController) RegisterTradeRoutes(rg *gin.RouterGroup) {
	tradeRoute := rg.Group("/trade")
	tradeRoute.POST("/create", TC.CreateTrade)
	tradeRoute.GET("/get/:id", TC.Get_SingleTrade)
	tradeRoute.GET("/getAll", TC.GetAllTrade)
	tradeRoute.PATCH("/update", TC.Update_SingleTradeContent)
	tradeRoute.DELETE("/delete/:id", TC.Delete_SingleTrade)
}
