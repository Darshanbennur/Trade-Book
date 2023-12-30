package handlers

import (
	"net/http"

	"github.com/Darshanbennur/Trade-Book/api/models"
	"github.com/Darshanbennur/Trade-Book/api/services"
	"github.com/gin-gonic/gin"
)

/*
Will contain Controller of static components like :
	- CRUD operations on Company Table
	-
*/

type CompanyController struct {
	CompanyService services.CompanyServices
}

func NewCompany(companyService services.CompanyServices) CompanyController {
	return CompanyController{
		CompanyService: companyService,
	}
}

func (CC *CompanyController) CreateCompany(ctx *gin.Context) {
	var company models.Companies
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := CC.CompanyService.CreateCompany(&company)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (CC *CompanyController) Get_SingleCompanyDetails(ctx *gin.Context) {
	companyName := ctx.Param("id")
	company, err := CC.CompanyService.Get_SingleCompanyDetails(&companyName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"instrument": company})
}

func (CC *CompanyController) GetAllCompanyDetails(ctx *gin.Context) {
	companies, err := CC.CompanyService.GetAllCompanyDetails()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"instrument": companies})
}

func (CC *CompanyController) Update_SingleCompany(ctx *gin.Context) {
	companyName := ctx.Param("id")
	var company models.Companies
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	erro := CC.CompanyService.Update_SingleCompany(&companyName, &company)
	if erro != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": erro.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updation Successful"})
}

func (CC *CompanyController) Delete_SingleCompany(ctx *gin.Context) {
	companyName := ctx.Param("id")
	err := CC.CompanyService.Delete_SingleCompany(&companyName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deletion Successful"})
}

func (CC *CompanyController) RegisterCompanyRoutes(rg *gin.RouterGroup) {
	companyRoute := rg.Group("/company")
	companyRoute.POST("/create", CC.CreateCompany)
	companyRoute.GET("/get/:id", CC.Get_SingleCompanyDetails)
	companyRoute.GET("/getAll", CC.GetAllCompanyDetails)
	companyRoute.PATCH("/update/:id", CC.Update_SingleCompany)
	companyRoute.DELETE("/delete/:id", CC.Delete_SingleCompany)
}
