package routes

import (
	"github.com/Sunchiii/champamker-service/controllers"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.RouterGroup) {
	route.POST("/login", controllers.Admin())
	route.POST("/register", controllers.Register())

}
