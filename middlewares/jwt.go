package middlewares

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId int, name string, email string, userRole string) string {
	var userClaims = jwtCustomClaims{
		userId, name, email, userRole,
		jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	resultJWT, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return resultJWT
}
