// internal/module/chat/router.go

package chat

import (
	"OpenTan/internal/global/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *ChatModule) InitRouter(r *gin.RouterGroup) {
	chatGroup := r.Group("/chat")
	{
		chatGroup.POST("/completions", func(c *gin.Context) {
			var req model.OpenAICompletionsRequest

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			handler := Completions(req)
			handler(c)
		})
	}
}
