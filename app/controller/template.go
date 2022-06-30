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

// AddTemplate
// @tags Template
// @summary Add New Template
// @router /template/add [post]
// @accept json
// @param data body model.TemplateRequest true "Template Information"
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Get uid or template wrong"
// @failure 500 "Add template failed"
func AddTemplate(c echo.Context) error {
	uid, ok := c.Get("uid").(uint)
	templateRequest := model.TemplateRequest{}
	var err error
	if err = c.Bind(&templateRequest); err == nil {
		err = c.Validate(&templateRequest)
	}
	if !ok || err != nil {
		logrus.Info(err)
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get uid or template wrong," + err.Error()})
	}
	template := model.Template{}
	err = copier.Copy(&template, templateRequest)
	template.UID = uid
	err = model.AddTemplate(&template)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "Add template failed"})
	}
	return c.JSON(http.StatusOK, response.Response{
		Msg: template.ID,
	})
}

// DeleteTemplate
// @tags Template
// @summary Delete Template
// @router /template/delete [post]
// @param tid query int true "template id"
// @produce json
// @accept json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Get tid wrong"
// @failure 500 "Delete template failed"
func DeleteTemplate(c echo.Context) error {
	var tid uint
	if err := echo.FormFieldBinder(c).MustUint("tid", &tid).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get tid wrong"})
	}
	if err := model.DeleteTemplate(tid); err != nil {
		logrus.Info(err)
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "Delete template failed"})
	}
	return c.JSON(http.StatusOK, response.Response{})
}

// GetAllTemplate
// @tags Template
// @summary Get All Template by uid
// @router /template/list [get]
// @accept json
// @produce json
// @success 200 {array} model.Template
// @failure 400 {object} response.Response10010 "Get uid wrong"
// @failure 500 "Get templates failed"
func GetAllTemplate(c echo.Context) error {
	uid, ok := c.Get("uid").(uint)
	if !ok {
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get uid wrong"})
	}
	Templates, err := model.GetAllTemplate(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "get templates failed"})
	}
	return c.JSON(http.StatusOK, response.Response{Msg: Templates})
}

// GetTemplate
// @tags Template
// @summary Get Template by id
// @router /template/get [get]
// @param tid query int true "template's id"
// @accept json
// @produce json
// @success 200 {object} model.GetTemplateInfoResp "template info"
// @failure 400 {object} response.Response10010 "Params error"
// @failure 404 {object} response.Response10020 "Invalid template ID"
// @failure 500 "Failed to get template info"
func GetTemplate(c echo.Context) error {
	var tid uint
	if err := echo.FormFieldBinder(c).MustUint("tid", &tid).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	result, err := model.QueryTemplateById(tid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Info(err)
		return c.JSON(http.StatusNotFound, response.Response{
			10020,
			"Invalid template ID",
		})
	} else if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Failed to get template info",
		})
	}

	ret := model.GetTemplateInfoResp{
		ID:         result.ID,
		Title:      result.Title,
		Subtasks:   result.Subtasks,
		EstimatedT: result.EstimatedT,
		Priority:   result.Priority,
	}
	return c.JSON(http.StatusOK, response.Response{
		Msg: &ret,
	})
}

// ChangeTemplate
// @tags Template
// @summary ChangeTodoInfo
// @description Change a template by id
// @router /template/change [post]
// @param tid query int true "template's id"
// @param data body model.ChangeTemplateReq true "template's new info"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "params error"
// @failure 500 "Failed to change template"
func ChangeTemplate(c echo.Context) error {
	var tid uint
	var newTemplate model.ChangeTemplateReq
	uid := c.Get("uid").(uint)
	if err := echo.FormFieldBinder(c).MustUint("tid", &tid).BindError(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	if err := c.Bind(&newTemplate); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			10010,
			"Params error: " + err.Error(),
		})
	}

	var templateInfo model.Template
	_ = copier.Copy(&templateInfo, newTemplate)
	templateInfo.ID = tid
	templateInfo.UID = uid

	if err := model.UpdateTemplateInfo(&templateInfo); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{
			Msg: "Failed to change template",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Msg: "OK",
	})
}
