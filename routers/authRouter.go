package routers

import (
	controller "github.com/NGODINHSANG/golang-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(imcomingRouters *gin.Engine) {
	imcomingRouters.POST("users/signup", controller.Signup())
	imcomingRouters.POST("users/login", controller.Login())
}
