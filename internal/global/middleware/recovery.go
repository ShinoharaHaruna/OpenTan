package middleware

import (
	"OpenTan/internal/global/response"
	"github.com/gin-gonic/gin"
	"log"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Recovery")
		defer response.Recovery(c)
		c.Next()
	}
}
