package controller

import (
	"TodoQueue/app/response"
	"TodoQueue/model"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewTodo(c echo.Context) error {
	newTodo := model.NewTodo()
	uid := c.Get("uid").(uint)
	if err := c.Bind(newTodo); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"params error",
		})
	}
	if newTodo.Priority <= 0 || newTodo.Priority > 4 {
		logrus.Error(errors.New("invalid priority"))
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"invalid priority",
		})
	}
	if newTodo.Title == "" {
		logrus.Error(errors.New("empty title in new todo"))
		return c.JSON(http.StatusBadRequest, response.Response{
			10002,
			"empty title",
		})
	}

	newTodo.UID = uid
	if err := model.CreateNewTodo(newTodo); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "failed to creat new todo",
		})
	}

	return nil
}

func GetTodoInfo(c echo.Context) error {

	return nil
}

func GetTodoList(c echo.Context) error {

	return nil
}

func ChangeTodoInfo(c echo.Context) error {

	return nil
}

func ArrangeTodo(c echo.Context) error {

	return nil
}

func DeleteTodo(c echo.Context) error {
	return nil
}

func GetDoneList(c echo.Context) error {
	return nil
}

func GetDoneInfo(c echo.Context) error {
	return nil
}

func DeleteDone(c echo.Context) error {
	return nil
}
