package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleAuthorization(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")
		authorized := false
		for _, role := range allowedRoles {
			if role == userRole {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
			c.Abort()
			return
		}
		c.Next()
	}
}
