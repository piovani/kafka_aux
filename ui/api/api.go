package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/piovani/kafka_aux/config"
)

type Api struct {
	engine *gin.Engine

	messageController InterfaceController
}

func NewApi() *Api {
	return &Api{
		engine: gin.Default(),
	}
}

func (a *Api) StartRest() error {
	a.GetControllers()
	a.GetRoutes()

	return a.engine.Run(fmt.Sprintf(":%d", config.Env.ApiRestPort))
}
