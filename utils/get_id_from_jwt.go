package utils

import (
	"e-complaint-api/constants"

	"github.com/labstack/echo/v4"
)

func GetIDFromJWT(c echo.Context) (int, error) {
	// get JwtToken from authorization header
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		return 0, constants.ErrUnauthorized
	}

	// Get JWT Token from Authorization Header
	jwtToken := GetToken(authorization)

	jwt_payload, err := DecodePayload(jwtToken)
	if err != nil {
		return 0, constants.ErrInternalServerError
	}

	// Get user id from jwt payload
	user_id, ok := jwt_payload["id"].(float64)
	if !ok {
		return 0, constants.ErrInternalServerError
	}

	return int(user_id), nil
}
