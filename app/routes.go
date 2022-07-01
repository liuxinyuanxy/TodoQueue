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
		AllowOrigins:     []string{"http://124.221.92.18:80"},
		AllowCredentials: true,
	}))
	//e.Use(echoMiddleware.BodyDumpWithConfig(DefaultBodyDumpConfig))
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

//var DefaultBodyDumpConfig = echoMiddleware.BodyDumpConfig{
//	Skipper: BodyDumpDefaultSkipper,
//	Handler: func(c echo.Context, reqBody []byte, resBody []byte) {
//		println("API请求结果拦截：", string(resBody), c.RealIP(), string(reqBody))
//		// 1、解析返回的json数据，判断接口执行成功或失败。如： {"code":"200","data":"test","msg":"请求成功"}
//		// 2、保存操作日志
//	},
//}
//
//// 排除文件，如果您的请求/响应有效负载非常大，例如文件上载/下载，需要进行排查。否则将影响响应时间
//func BodyDumpDefaultSkipper(c echo.Context) bool {
//	if strings.Contains(c.Path(), "/api/files/") {
//		return true
//	}
//	return false
//}
