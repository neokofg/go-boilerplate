package helpers

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/infrastructure/ent"
	"net/http"
)

func HandleServerErr(c *gin.Context, err error, msg string) bool {
	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
			return true
		}
		if ent.IsConstraintError(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "constraint violation"})
			return true
		}
		if msg == "" {
			msg = "Internal server error"
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return true
	}
	return false
}
