package controller

import (
	"TodoQueue/app/response"
	"TodoQueue/model"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get uid or template wrong"})
	}
	template := model.Template{}
	err = copier.Copy(&template, templateRequest)
	template.UID = uid
	err = model.AddTemplate(&template)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "Add template failed"})
	}
	return c.JSON(http.StatusOK, response.Response{})
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

// GetTemplate
// @tags Template
// @summary Delete Template
// @router /template/get [get]
// @produce json
// @success 200 {array} model.Template
// @failure 400 {object} response.Response10010 "Get uid wrong"
// @failure 500 "Get template failed"
func GetTemplate(c echo.Context) error {
	uid, ok := c.Get("uid").(uint)
	if !ok {
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get uid wrong"})
	}
	Templates, err := model.GetAllTemplate(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "get template failed"})
	}
	return c.JSON(http.StatusOK, response.Response{Msg: Templates})
}
