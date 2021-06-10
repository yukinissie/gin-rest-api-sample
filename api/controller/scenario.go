package controller

import (
	"main/model"
	service "main/service/scenario"
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
	setScenarioService := service.SetScenarioService{}
	err = setScenarioService.Excute(&scenario)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func ScenarioList(c *gin.Context) {
	getScenarioService := service.GetScenarioService{}
	ScenarioLists, err := getScenarioService.Excute()
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    ScenarioLists,
	})
}
