package routes

import (
	"github.com/Sunchiii/champamker-service/controllers"
	"github.com/gin-gonic/gin"
)

func Initial(route *gin.Engine) {
	route.GET("ping", controllers.Initial())
}
