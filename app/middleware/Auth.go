package middleware

import (
	"github.com/labstack/echo/v4"
)

// Auth
// check the token and set uid
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("uid", uint(1))
		return next(c)
		//	cookie, err := c.Cookie("token")
		//	if err == nil {
		//		tokenString := cookie.Value
		//		var token *jwt.Token
		//		claims := &jwt.StandardClaims{}
		//		token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		//			return []byte("thisisAkeyqwq"), nil
		//		})
		//		if err == nil && token.Valid {
		//			uid, _ := strconv.Atoi(claims.Subject)
		//			c.Set("uid", uint(uid))
		//			return next(c)
		//		}
		//	}
		//	return c.JSON(http.StatusForbidden, response.Response{Code: 10030, Msg: "Please login first"})
	}
}
