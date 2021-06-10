package controller

import (
	"main/model"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScenarioAdd(c *gin.Context) {
	scenario := model.Scenario{}
	err := c.Bind(&scenario)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	scenarioService := service.ScenarioService{}
	err = scenarioService.SetScenario(&scenario)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func ScenarioList(c *gin.Context) {
	scenarioService := service.ScenarioService{}
	ScenarioLists, err := scenarioService.GetScenarioList()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    ScenarioLists,
	})
}
