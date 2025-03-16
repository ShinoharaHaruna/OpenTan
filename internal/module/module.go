package module

import (
	"OpenTan/internal/module/chat"
	"OpenTan/internal/module/models"
	"OpenTan/internal/module/ping"
	"github.com/gin-gonic/gin"
)

type Module interface {
	GetName() string
	Init()
	InitRouter(r *gin.RouterGroup)
}

var Modules []Module

func registerModule(m []Module) {
	Modules = append(Modules, m...)
}

func init() {
	registerModule([]Module{
		&ping.ModulePing{},
		&models.ShowModels{},
		&chat.ChatModule{},
	})
}
