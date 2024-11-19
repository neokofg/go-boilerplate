package handlers

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/helpers"
	"go-boilerplate/infrastructure/ent"
	"net/http"
)

type UserHandler struct {
	client *ent.Client
}

func NewUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{
		client: client,
	}
}

func (h *UserHandler) GetSelf(c *gin.Context) {
	user, err := helpers.GetUser(c)
	helpers.HandleServerErr(c, err, "user get failed")

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
	})
}
