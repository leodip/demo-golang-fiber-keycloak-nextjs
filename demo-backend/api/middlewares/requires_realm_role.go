package middlewares

import (
	"demo-backend/shared/enums"
	"demo-backend/shared/jwt"

	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
)

func NewRequiresRealmRole(role string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		claims := ctx.Value(enums.ContextKeyClaims).(golangJwt.MapClaims)
		jwtHelper := jwt.NewJwtHelper(claims)
		if !jwtHelper.IsUserInRealmRole(role) {
			return c.Status(fiber.StatusUnauthorized).SendString("role authorization failed")
		}
		return c.Next()
	}
}
