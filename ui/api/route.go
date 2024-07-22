package api

import "github.com/piovani/kafka_aux/ui/api/controller"

func (a *Api) GetControllers() {
	a.messageController = controller.NewMessageController()
}

func (a *Api) GetRoutes() {
	a.engine.GET("/messages", a.messageController.List)
	a.engine.POST("/messages", a.messageController.Create)
}
