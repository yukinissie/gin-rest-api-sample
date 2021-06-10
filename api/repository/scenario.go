package repository

import (
	"errors"
	"main/model"
	"main/repository/redis"
)

func SetScenario(model *model.Scenario) error {
	switch PERSISTENCE_MODE {
	case "redis":
		return redis.Set(SCENARIO_PATH, model)
	default:
		return errors.New("invalid persistence mode")
	}
}

func GetAllScenario() ([]model.Scenario, error) {
	switch PERSISTENCE_MODE {
	case "redis":
		return redis.Get(SCENARIO_PATH), nil
	default:
		return nil, errors.New("invalid persistence mode")
	}
}
