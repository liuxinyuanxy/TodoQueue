package controller

import (
	"TodoQueue/app/response"
	"TodoQueue/model"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

// StartProgress
// @tags Progress
// @summary StartProgress
// @description Start a todo by id
// @router /progress/start [post]
// @param id query int true "todo's id"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Params error"
// @failure 404 {object} response.Response10020 "Invalid ID"
// @failure 500 "database error"
func StartProgress(c echo.Context) error {
	var id uint
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	err := model.StartTodoById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, response.Response{
			10020,
			"Invalid ID",
		})
	}
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Database error",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}

// SuspendProgress
// @tags Progress
// @summary SuspendProgress
// @description Suspend the todo in progress
// @router /progress/suspend [post]
// @param id query int true "Id of todo which in progess"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Params error"
// @failure 404 {object} response.Response10020 "Invalid ID"
// @failure 406 {object} response.Response10011 "Todo doesn't in progress"
// @failure 500 "database error"
func SuspendProgress(c echo.Context) error {
	var id uint
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	err := model.SuspendTodoById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, response.Response{
			10020,
			"Invalid ID",
		})
	}
	if errors.Is(err, gorm.ErrInvalidData) {
		return c.JSON(http.StatusNotAcceptable, response.Response{
			Code: 10011,
			Msg:  "Todo doesn't in process",
		})
	}
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Database error",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}

// FinishProgress
// @tags Progress
// @summary FinishProgress
// @description Finish a todo by id, move it to done
// @router /progress/finish [post]
// @param id query int true "todo's id"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Params error"
// @failure 404 {object} response.Response10020 "Invalid ID"
// @failure 500 "database error"
func FinishProgress(c echo.Context) error {
	var id uint
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	err := model.MoveTodo2Done(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, response.Response{
			10020,
			"Invalid ID",
		})
	}
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Database error",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}
