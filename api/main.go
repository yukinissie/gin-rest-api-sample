package main

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	apiEngine := engine.Group("/api")
	{
		apiEngine.POST("/scenario", controller.ScenarioAdd)
		apiEngine.GET("/scenario", controller.ScenarioList)
	}
	engine.Run(":3000")
}
