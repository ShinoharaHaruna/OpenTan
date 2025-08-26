package server

import (
	"OpenTan/config"
	"OpenTan/internal/global"
	"OpenTan/internal/global/logger"
	"OpenTan/internal/global/middleware"
	"OpenTan/internal/module"
	"OpenTan/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

var log *slog.Logger

func Init() {
	log = logger.New("Server")

	for _, m := range module.Modules {
		log.Info(fmt.Sprintf("Init Module: %s", m.GetName()))
		m.Init()
	}
}

func Run() {
	c := config.Get()
	gin.SetMode(string(c.Mode))
	r := gin.New()

	switch c.Mode {
	case config.ModeRelease:
		r.Use(middleware.Logger(logger.Get()))
	case config.ModeDebug:
		r.Use(gin.Logger())
	}

	r.Use(middleware.Recovery())

	for _, m := range module.Modules {
		log.Info(fmt.Sprintf("Init Router: %s", m.GetName()))
		m.InitRouter(r.Group("/" + c.Prefix))
	}

	// Try refreshing token every launch of the server
	updated := global.TryRefresh()
	if updated && config.IsDebug() {
		log.Info(fmt.Sprintf("Token refreshed: %s", c.API_KEY))
	}

	err := r.Run(c.Host + ":" + c.Port)
	utils.PanicOnErr(err)
}
