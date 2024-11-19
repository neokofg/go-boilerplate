package middleware

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/ent"
	"go-boilerplate/infrastructure/ent/user"
	"go-boilerplate/pkg/jwt"
	"os"
	"strings"
)

type AuthMiddleware struct {
	client *ent.Client
}

func NewAuthMiddleware(client *ent.Client) *AuthMiddleware {
	return &AuthMiddleware{
		client: client,
	}
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.ValidateToken(tokenString, os.Getenv("JWT_SECRET"))
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		u, err := m.client.User.Query().
			Where(user.ID(int(claims.UserID))).
			Only(c.Request.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				c.JSON(401, gin.H{"error": "user not found"})
			} else {
				c.JSON(500, gin.H{"error": "database error"})
			}
			c.Abort()
			return
		}

		c.Set("user", u)

		c.Next()
	}
}
