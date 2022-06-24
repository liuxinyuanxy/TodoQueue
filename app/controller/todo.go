package controller

import (
	"TodoQueue/app/response"
	"TodoQueue/model"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

// NewTodo
// @tags Todo
// @summary NewTodo
// @description create a new todo
// @router /todo/new [post]
// @param data body model.NewTodoReq true "new todo info"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "params error"
// @failure 400 {object} response.Response10010 "invalid priority"
// @failure 400 {object} response.Response10002 "empty title"
// @failure 500 "database error"
func NewTodo(c echo.Context) error {
	newTodo := model.NewTodo()
	uid := c.Get("uid").(uint)
	if err := c.Bind(newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}
	if newTodo.Priority <= 0 || newTodo.Priority > 4 {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Invalid priority",
		})
	}
	if newTodo.Title == "" {
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

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}

// GetTodoInfo
// @tags Todo
// @summary GetTodoInfo
// @description Get a todo's info by id
// @router /todo/get [get]
// @param id query int true "todo's id"
// @accept json
// @produce json
// @success 200 {object} model.Todo "todo info"
// @failure 400 {object} response.Response10010 "params error"
// @failure 404 {object} response.Response10020 "Invalid todo ID"
// @failure 500 "database error"
func GetTodoInfo(c echo.Context) error {
	var id uint
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Praram error: " + err.Error(),
		})
	}

	result, err := model.QueryTodoById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Info(err)
		return c.JSON(http.StatusNotFound, response.Response{
			10020,
			"Invalid todo ID",
		})
	} else if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Database error",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: &result,
	})
}

// GetTodoList
// @tags Todo
// @summary GetTodoList
// @description Get a todolist by uid
// @router /todo/list [get]
// @accept json
// @produce json
// @success 200 {array}  model.Todo "Todo info list"
// @failure 500 "database error"
func GetTodoList(c echo.Context) error {
	uid := c.Get("uid").(uint)

	result, err := model.QueryOrderedTodoListByUID(uid)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Database error",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: &result,
	})
}

// ChangeTodoInfo
// @tags Todo
// @summary ChangeTodoInfo
// @description Change a todo's info by id
// @router /todo/change [post]
// @param id query int true "todo's id"
// @param data body model.ChangeTodoInfoReq true "todo's new info"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "params error"
// @failure 500 "database error"
func ChangeTodoInfo(c echo.Context) error {
	var id uint
	var newTodo model.ChangeTodoInfoReq
	uid := c.Get("uid").(uint)
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	if err := c.Bind(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	var todoInfo model.Todo
	_ = copier.Copy(&todoInfo, newTodo)
	todoInfo.ID = id
	todoInfo.UID = uid

	if err := model.UpdateTodoInfo(&todoInfo); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Failed to change todo info",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}

// DeleteTodo
// @tags Todo
// @summary DeleteTodo
// @description Delete a todo by id
// @router /todo/delete [post]
// @param id query int true "todo's id"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "params error"
// @failure 500 "database error"
func DeleteTodo(c echo.Context) error {
	var id uint
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	if err := model.DeleteTodoById(id); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "failed to delete todo",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}

// GetDoneList
// @tags Todo
// @summary GetDoneList
// @description Get a done list by uid
// @router /todo/list/done [get]
// @accept json
// @produce json
// @success 200 {array}  model.TodoDone "Done info list"
// @failure 500 "database error"
func GetDoneList(c echo.Context) error {
	uid := c.Get("uid").(uint)

	result, err := model.QueryDoneListByUID(uid)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Database error",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: &result,
	})
}

// GetDoneInfo
// @tags Todo
// @summary GetDoneInfo
// @description Get a done's info by id
// @router /todo/get/done [get]
// @param id query int true "done's id"
// @accept json
// @produce json
// @success 200 {object} model.TodoDone "done info"
// @failure 400 {object} response.Response10010 "params error"
// @failure 404 {object} response.Response10020 "Invalid done ID"
// @failure 500 "database error"
func GetDoneInfo(c echo.Context) error {
	var id uint
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Praram error: " + err.Error(),
		})
	}

	result, err := model.QueryDoneById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Info(err)
		return c.JSON(http.StatusNotFound, response.Response{
			10020,
			"Invalid done ID",
		})
	} else if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Database error",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: &result,
	})
}

// DeleteDone
// @tags Todo
// @summary DeleteDone
// @description Delete a done by id
// @router /todo/delete/done [post]
// @param id query int true "done's id"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "params error"
// @failure 500 "database error"
func DeleteDone(c echo.Context) error {
	var id uint
	if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	if err := model.DeleteDoneById(id); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "failed to delete done",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}
