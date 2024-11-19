package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/ent"
)

func GetUser(c *gin.Context) (*ent.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("user not found in context")
	}

	u, ok := user.(*ent.User)
	if !ok {
		return nil, fmt.Errorf("invalid user type in context")
	}

	return u, nil
}
