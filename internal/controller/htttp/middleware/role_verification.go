package middleware

import (
	"net/http"

	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"
)


const (
	roleModerator = "moderator"
	rolePvzEmployee = "pvz_employee"
)

func (m *Middleware) ModeratorVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, ok := c.Get("userRole")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"message": errorsx.ErrNoPermission.Error()})
			c.Abort()
			return
		}
		if userRole != roleModerator {
			c.JSON(http.StatusForbidden, gin.H{"message": errorsx.ErrNoPermission.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (m *Middleware) PvzEmployeeVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, ok := c.Get("userRole")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"message": errorsx.ErrNoPermission.Error()})
			c.Abort()
			return
		}
		if userRole != rolePvzEmployee {
			c.JSON(http.StatusForbidden, gin.H{"message": errorsx.ErrNoPermission.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
