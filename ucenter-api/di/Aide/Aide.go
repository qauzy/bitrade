package Aide

import (
	"bitrade/ucenter-api/controller"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*gin.Context)

func KeyWords(controller *controller.AideController) HandlerFunc {
	return func(ctx *gin.Context) {
		result := controller.KeyWords(ctx)
		ctx.JSON(200, result)
	}

}
