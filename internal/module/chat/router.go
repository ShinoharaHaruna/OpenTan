// internal/module/chat/router.go

package chat

import (
	"OpenTan/internal/global/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var mutex sync.Mutex

func (m *ChatModule) InitRouter(r *gin.RouterGroup) {
	chatGroup := r.Group("/chat")
	{
		chatGroup.POST("/completions", func(c *gin.Context) {
			mutex.Lock()
			// 由于 AI 端的消费能力实在有限，干脆做串行了
			defer mutex.Unlock()

			var req model.OpenAICompletionsRequest

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			//log.Printf("Request size: %d Bytes\n", unsafe.Sizeof(req))
			handler := Completions(req)
			handler(c)
		})
	}
}
