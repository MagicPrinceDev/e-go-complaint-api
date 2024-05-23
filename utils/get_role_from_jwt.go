package utils

import (
	"e-complaint-api/constants"

	"github.com/labstack/echo/v4"
)

func GetRoleFromJWT(c echo.Context) (string, error) {
	// get JwtToken from authorization header
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		return "", constants.ErrUnauthorized
	}

	// Get JWT Token from Authorization Header
	jwtToken := GetToken(authorization)

	jwt_payload, err := DecodePayload(jwtToken)
	if err != nil {
		return "", constants.ErrInternalServerError
	}

	// Get user role from jwt payload
	role, ok := jwt_payload["role"].(string)
	if !ok {
		return "", constants.ErrInternalServerError
	}

	return role, nil
}
