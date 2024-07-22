package api

import "github.com/gin-gonic/gin"

type InterfaceController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
}
