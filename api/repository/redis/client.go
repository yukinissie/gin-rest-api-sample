package redis

import (
	"encoding/json"
	"log"
	"main/model"
	"main/persistence/redis"
	"os"

	"github.com/pkg/errors"
)

func Set(savePath string, model *model.Scenario) error {
	redisPath := os.Getenv("REDIS_PATH")
	log.Println(redisPath)
	client, err := redis.New(redisPath)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 保存用オブジェクトの作成
	saveObject, err := json.Marshal(model)
	if err != nil {
		log.Printf("cannot marshal struct: %v\n", err)
		return err
	}
	err = client.SAdd(savePath, saveObject).Err()
	if err != nil {
		return errors.Wrap(err, "Failed to save item")
	}
	return nil
}

func Get(savePath string) []model.Scenario {
	redisPath := os.Getenv("REDIS_PATH")
	log.Println(redisPath)
	client, err := redis.New(redisPath)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	responseObject, _, err := client.SScan(savePath, 0, "", 100).Result()
	if err != nil {
		log.Fatal(err)
	}

	responseBody := make([]model.Scenario, 0)
	for i := 0; i < len(responseObject); i++ {
		scenario := new(model.Scenario)
		json.Unmarshal([]byte(responseObject[i]), &scenario)
		responseBody = append(responseBody, *scenario)
	}
	return responseBody
}
