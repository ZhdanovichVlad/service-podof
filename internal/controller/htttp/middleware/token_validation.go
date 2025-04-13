package middleware

import (
	"net/http"
	"strings"


	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"

)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errorsx.ErrAuthHeaderIsEmpty.Error()})
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)


		claims, err := m.tokenValidator.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		
			c.Set("userUUID", claims.UserID)
			c.Set("userRole", claims.Role)
		c.Next()
	}
}

