package app

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-svc-tpl/app/controller"
	"go-svc-tpl/app/middleware"
)

func addRoutes() {
	api := e.Group("api")

	api.GET("/doc/*", echoSwagger.WrapHandler)

	user := api.Group("user")
	user.GET("/logout", controller.LogOut)
	user.POST("/login", controller.LogIn)
	user.POST("/register", controller.SignIn)
	userChange := user.Group("change", middleware.Auth)
	userChange.POST("/passwd", controller.ChangePassword)
	userChange.POST("/name", controller.ChangeNickname)
}
