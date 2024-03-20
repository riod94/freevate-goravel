package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Index(ctx http.Context) http.Response {
	var users []models.User
	facades.Orm().Query().Find(&users)
	return ctx.Response().Success().Json(http.Json{
		"result": users,
	})
}

func (r *UserController) Store(ctx http.Context) http.Response {
	user := models.User{Name: "tom", Email: "tom@qq.com", Password: "123456"}
	result := facades.Orm().Query().Create(&user)
	return ctx.Response().Success().Json(http.Json{
		"result": result,
	})
}

func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var user models.User
	facades.Orm().Query().Find(&user, id)
	return ctx.Response().Success().Json(http.Json{
		"result": user,
	})
}

func (r *UserController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var user models.User
	facades.Orm().Query().Find(&user, id)
	user.Name = ctx.Request().Input("name", user.Name)
	user.Email = ctx.Request().Input("email", user.Email)
	password, _ := facades.Hash().Make(ctx.Request().Input("password", user.Password))
	user.Password = password
	facades.Orm().Query().Save(&user)
	return ctx.Response().Success().Json(http.Json{
		"result": user,
	})
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var user models.User
	facades.Orm().Query().Find(&user, id)
	facades.Orm().Query().Delete(&user)
	return ctx.Response().Success().Json(http.Json{
		"result": user,
	})
}
