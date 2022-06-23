package app

import (
	"TodoQueue/app/controller"
	"TodoQueue/app/middleware"
	"TodoQueue/app/response"
	"TodoQueue/model"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func ping(c echo.Context) error {

	newDone := model.TodoDone{
		ID:    2,
		UID:   1,
		Title: "asf",
	}
	model.CreateDone(&newDone)

	c.JSON(http.StatusOK, response.Response{
		Msg: "pong!",
	})
	return nil
}

func addRoutes() {
	api := e.Group("api")

	api.GET("/doc/*", echoSwagger.WrapHandler)
	api.POST("/ping", ping)

	user := api.Group("/user")
	{
		user.GET("/logout", controller.LogOut)
		user.POST("/login", controller.LogIn)
		user.POST("/register", controller.SignIn)
		userChange := user.Group("/change", middleware.Auth)
		userChange.POST("/passwd", controller.ChangePassword)
		userChange.POST("/name", controller.ChangeNickname)
	}

	template := api.Group("/template", middleware.Auth)
	{
		template.GET("/get", controller.GetTemplate)
		template.POST("/delete", controller.DeleteTemplate)
		template.POST("/add", controller.AddTemplate)
	}

	todo := api.Group("/todo", middleware.Auth, middleware.CheckOwner)
	{
		todo.POST("/new", controller.NewTodo)
		todo.POST("/delete", controller.DeleteTodo)
		todo.POST("/change", controller.ChangeTodoInfo)
		todo.POST("/delete/done", controller.DeleteDone)
		todo.GET("/get", controller.GetTodoInfo)
		todo.GET("/list", controller.GetTodoList)
		todo.GET("/get/done", controller.GetDoneInfo)
		todo.GET("/list/done", controller.GetDoneList)
	}

	progress := api.Group("/progress", middleware.Auth)
	{
		progress.POST("/tp", nil)
	}
}
