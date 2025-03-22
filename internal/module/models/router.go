package models

import (
	"OpenTan/internal/global"
	"github.com/gin-gonic/gin"
)

func (p *ShowModels) InitRouter(r *gin.RouterGroup) {
	r.GET("/models", global.GetModels())
}
