package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type BankController struct {
	//Dependent services
}

func NewBankController() *BankController {
	return &BankController{
		//Inject services
	}
}

func (r *BankController) Index(ctx http.Context) http.Response {
	banks := []models.Bank{}
	facades.Orm().Query().Get(&banks)
	return ctx.Response().Success().Json(http.Json{
		"result": banks,
	})
}

func (r *BankController) Store(ctx http.Context) http.Response {
	bank := models.Bank{
		Name:     ctx.Request().Input("name"),
		Code:     ctx.Request().Input("code"),
		Status:   ctx.Request().Input("status"),
		Category: ctx.Request().Input("category"),
		Link:     ctx.Request().Input("link"),
		Contact:  ctx.Request().Input("contact"),
	}

	result := facades.Orm().Query().Create(&bank)

	return ctx.Response().Success().Json(http.Json{
		"result": result,
	})
}

func (r *BankController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	bank := models.Bank{}
	facades.Orm().Query().Where("id", id).OrWhere("code", id).Find(&bank)
	return ctx.Response().Success().Json(http.Json{
		"result": bank,
	})
}

func (r *BankController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	bank := models.Bank{}

	facades.Orm().Query().Where("id", id).OrWhere("code", id).Find(&bank)

	bank.Name = ctx.Request().Input("name", bank.Name)
	bank.Code = ctx.Request().Input("code", bank.Code)
	bank.Status = ctx.Request().Input("status", bank.Status)
	bank.Category = ctx.Request().Input("category", bank.Category)
	bank.Link = ctx.Request().Input("link", bank.Link)
	bank.Contact = ctx.Request().Input("contact", bank.Contact)
	facades.Orm().Query().Save(&bank)

	return ctx.Response().Success().Json(http.Json{
		"result": bank,
	})
}

func (r *BankController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	bank := models.Bank{}
	facades.Orm().Query().Where("id", id).OrWhere("code", id).Find(&bank)
	facades.Orm().Query().Delete(&bank)
	return ctx.Response().Success().Json(http.Json{
		"result": bank,
	})
}
