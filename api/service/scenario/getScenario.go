package service

import (
	"main/model"
	"main/repository"
)

type GetScenarioService struct{}

func (GetScenarioService) Excute() ([]model.Scenario, error) {
	scenarioList, err := repository.GetAllScenario()
	return scenarioList, err
}
