package main

import (
	"github.com/Sunchiii/champamker-service/configs"
	"github.com/Sunchiii/champamker-service/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//log server status when begin

	//set release mode for production
	gin.SetMode(gin.ReleaseMode)

	//initial route
	route := gin.Default()
	route.Use(cors.Default())

	//connect database
	configs.ConnectDatebase()
	public := route.Group("/public")
	member := route.Group("/member")
	routes.Initial(route)
	routes.Admin(public)
	routes.Member(member)

	//log status when finish
	route.Run()
	
}
