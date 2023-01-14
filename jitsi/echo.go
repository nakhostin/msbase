package jitsi

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (sdk JitsiSDK) echoRoutes(GP *echo.Group) {
	GP.POST("/meet", sdk.createJitsiMeet)
}

//  api functions

// createJitsiMeet , create jitsi room and generate room link and jwt token
// meet link to other users
func (sdk JitsiSDK) createJitsiMeet(c echo.Context) error {

	form := struct {
		Token string `form:"token" json:"token"`
	}{}

	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	claims, err := readToken(form.Token, sdk.config.Echo.JWTSecret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	meetURL, err := sdk.CreateMeetWithClaims(claims)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	// return c.Redirect(http.StatusFound, meetURL)
	return c.JSON(http.StatusOK, echo.Map{"claims": claims, "meetURL": meetURL})
}

func readToken(tokenString, secret string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
