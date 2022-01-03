package api

import (
	"server/api/controller"
	"server/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.Use(middleware.ResponseHanlder())
	r.POST("/order", controller.ShoppingCartOrder)
	r.POST("/query", controller.ShoppingCartQuery)
	r.POST("/update", controller.ShoppingCartUpdate)
}
