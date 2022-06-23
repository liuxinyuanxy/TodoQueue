package middleware

import (
	"TodoQueue/app/response"
	"TodoQueue/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CheckOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id uint
		uid := c.Get("uid").(uint)
		if err := echo.FormFieldBinder(c).MustUint("id", &id).BindError(); err != nil {
			return next(c)
		}

		flag, err := model.CertificateTodo(id, uid)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, response.Response{
				Msg: "Database error",
			})
		}

		if flag {
			return next(c)
		}

		return c.JSON(http.StatusForbidden, response.Response{
			10030,
			"Todo doesn't belong to user",
		})
	}
}
