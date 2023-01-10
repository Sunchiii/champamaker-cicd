package routes

import (
	"github.com/Sunchiii/champamker-service/controllers"
	"github.com/gin-gonic/gin"
)

func Member(route *gin.RouterGroup) {
	admin := route.Group("/admin")
	{
		admin.POST("/", controllers.CreateMember())
		admin.PUT("/:id", controllers.Update())
		admin.DELETE("/:id", controllers.DeleteMember())
	}
	user := route.Group("/user")
	{
		user.GET("/", controllers.GetAllUser())
		user.GET("/:id", controllers.GetById())
	}

}
