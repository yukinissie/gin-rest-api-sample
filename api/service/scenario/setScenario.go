package service

import (
	"main/model"
	"main/repository"
)

type SetScenarioService struct{}

func (SetScenarioService) Excute(scenario *model.Scenario) error {
	return repository.SetScenario(scenario)
}
