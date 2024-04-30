package routers

import (
	controller "github.com/NGODINHSANG/golang-jwt/controllers"

	"github.com/NGODINHSANG/golang-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRouters *gin.Engine) {
	incomingRouters.Use(middleware.Authentication())
	incomingRouters.GET("/users", controller.GetUsers())
	incomingRouters.GET("/users/:user_id", controller.GetUser())
}
