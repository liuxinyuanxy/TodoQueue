package app

import (
	"TodoQueue/app/controller"
	"TodoQueue/app/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func ping(c echo.Context) error {
	logrus.Info(c.RealIP())
	return c.String(http.StatusOK, "pong from "+c.RealIP())
}

func addRoutes() {

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://124.221.92.18:9404", "http://localhost:9404"},
		AllowCredentials: true,
	}))
	api := e.Group("api")

	api.GET("/doc/*", echoSwagger.WrapHandler)
	api.POST("/ping", ping)

	user := api.Group("/user")
	{
		user.GET("/logout", controller.LogOut)
		user.POST("/login", controller.LogIn)
		user.POST("/register", controller.SignIn)
		user.GET("/get", controller.GetUserInfo, middleware.Auth)
		userChange := user.Group("/change", middleware.Auth)
		userChange.POST("/passwd", controller.ChangePassword)
		userChange.POST("/name", controller.ChangeNickname)
	}

	template := api.Group("/template", middleware.Auth)
	{
		template.GET("/get", controller.GetTemplate)
		template.GET("/list", controller.GetAllTemplate)
		template.POST("/delete", controller.DeleteTemplate)
		template.POST("/add", controller.AddTemplate)
		template.POST("/change", controller.ChangeTemplate)
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

	progress := api.Group("/progress", middleware.Auth, middleware.CheckOwner)
	{
		progress.POST("/start", controller.StartProgress)
		progress.POST("/suspend", controller.SuspendProgress)
		progress.POST("/finish", controller.FinishProgress)
		progress.GET("/get", controller.GetProgress)

	}
}
