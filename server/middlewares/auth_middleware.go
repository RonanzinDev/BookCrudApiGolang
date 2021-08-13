package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ronanzindev/book-api-golang/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJwtServices().ValidateToken(token) {
			c.AbortWithStatus(401)
		}

	}
}