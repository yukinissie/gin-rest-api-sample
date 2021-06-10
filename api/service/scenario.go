package service

import (
	"main/model"
	"main/repository"
)

type ScenarioService struct{}

func (ScenarioService) SetScenario(scenario *model.Scenario) error {
	return repository.SetScenario(scenario)
}

func (ScenarioService) GetScenarioList() ([]model.Scenario, error) {
	scenarioList, err := repository.GetAllScenario()
	return scenarioList, err
}
