package middlewares

import (
	"e-complaint-api/constants"
	"e-complaint-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, error := utils.GetRoleFromJWT(c)
		if error != nil || role != "super_admin" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": constants.ErrUnauthorized.Error(),
			})
		}

		return next(c)
	}
}
