package controller

import (
	"TodoQueue/app/response"
	"TodoQueue/model"
	"TodoQueue/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// SignIn
// @tags User
// @summary Sign in
// @description a new user signs in
// @router /user/register [post]
// @param email query string true "user's email"
// @param passwd query string true "user's password"
// @param name query string true "user's nickname"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "get name or passwd wrong"
// @failure 202 {object} response.Response10001 "email has been used"
// @failure 500 {object} "Create user failed"
func SignIn(c echo.Context) error {
	// first, we get the name and password
	var email, passwd, name string
	err := echo.FormFieldBinder(c).MustString("email", &email).MustString("passwd", &passwd).MustString("name", &name).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get name or passwd wrong"})
	}
	// then, we check if there is a user with same name
	_, err = model.QueryUserByEmail(email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusAccepted, response.Response{Code: 10001, Msg: "email has been used"})
	}
	// finally, we create a new user
	password := []byte(passwd)
	err = model.CreateUser(&model.User{Email: email, Password: password, Nickname: name})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "Create user failed"})
	}
	return c.JSON(http.StatusOK, response.Response{Msg: "Created user " + name})
}

// LogIn
// @tags User
// @summary Log in
// @description a user logs in
// @router /user/login [post]
// @param email query string true "user's email"
// @param passwd query string true "user's password"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Get name or passwd wrong"
// @failure 404 {object} response.Response10020 "User not found"
// @failure 403 {object} response.Response10030 "Wrong password"
// @failure 500 {object} "Can not generate token"
func LogIn(c echo.Context) error {
	// first, we get name and password
	var email, passwd string
	err := echo.FormFieldBinder(c).MustString("email", &email).MustString("passwd", &passwd).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get name or passwd wrong"})
	}
	password := []byte(passwd)
	// then, we check whether the user exists and whether the password is correct
	dbUser, err := model.QueryUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, response.Response{Code: 10020, Msg: "User not found"})
	}
	if bcrypt.CompareHashAndPassword(dbUser.Password, password) == bcrypt.ErrMismatchedHashAndPassword {
		return c.JSON(http.StatusForbidden, response.Response{Code: 10030, Msg: "wrong password"})
	}
	// if success, we set the token so that the user needn't log in every time
	tokenString, expireAt, err := utils.GenerateJwt(dbUser.ID)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "Can not generate token"})
	}
	cookies := &http.Cookie{Name: "token", Value: tokenString, Expires: expireAt, Path: "/"}
	c.SetCookie(cookies)
	return c.JSON(http.StatusOK, response.Response{Msg: "set token successfully"})
}

// LogOut
// @tags User
// @summary Log out
// @description a user logs out
// @router /user/logout [get]
func LogOut(c echo.Context) error {
	// we can easily log out by erase the cookie
	c.SetCookie(&http.Cookie{Expires: time.Now().Add(-time.Minute)})
	return nil
}

// ChangePassword
// @tags User
// @summary Change Password
// @router /user/change/passwd [post]
// @param passwd query string true "user's old password"
// @param newPasswd query string true "user's new password"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Get uid or passwd wrong"
// @failure 403 {object} response.Response10031 "User invalid, please log out and log in again"
// @failure 403 {object} response.Response10030 "Wrong password"
// @failure 500 "Change password failed"
func ChangePassword(c echo.Context) error {
	uid, ok := c.Get("uid").(uint)
	var passwd, newPasswd string
	err := echo.FormFieldBinder(c).MustString("newPasswd", &newPasswd).MustString("passwd", &passwd).BindError()
	if !ok || err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get uid or passwd wrong"})
	}
	dbUser, err := model.QueryUserByUid(uid)
	if err != nil {
		return c.JSON(http.StatusForbidden, response.Response{Code: 10031, Msg: "User invalid, please log out and log in again"})
	}
	password := []byte(passwd)
	if bcrypt.CompareHashAndPassword(dbUser.Password, password) == bcrypt.ErrMismatchedHashAndPassword {
		return c.JSON(http.StatusForbidden, response.Response{Code: 10030, Msg: "wrong password"})
	}
	err = model.UpdateUserPassword(uid, []byte(newPasswd))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "Change password failed"})
	}
	cookies, err := utils.GenerateCookie(uid)
	if err == nil {
		c.SetCookie(cookies)
	}
	return c.JSON(http.StatusOK, response.Response{Msg: "Change password success"})
}

// ChangeNickname
// @tags User
// @summary Change Nickname
// @router /user/change/name [post]
// @param name query string true "user's new nickname"
// @accept json
// @produce json
// @success 200 "OK"
// @failure 400 {object} response.Response10010 "Get uid or passwd wrong"
// @failure 500 "Change password failed"
func ChangeNickname(c echo.Context) error {
	uid, ok := c.Get("uid").(uint)
	var nickname string
	err := echo.FormFieldBinder(c).MustString("name", &nickname).BindError()
	if !ok || err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Code: 10010, Msg: "get uid or nickname wrong"})
	}
	err = model.UpdateUserNickname(uid, nickname)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Msg: "Change nickname failed"})
	}
	return c.JSON(http.StatusOK, response.Response{Msg: "Change nickname success"})
}
