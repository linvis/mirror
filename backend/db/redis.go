package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v7"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Db       int
}

var (
	r *redis.Client
)

func loadRedisConfig(file string) (*RedisConfig, error) {
	var config RedisConfig

	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return &config, nil
}

func InitRedis(path string) {

	config, err := loadRedisConfig(path)
	if err != nil {
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Db, // use default DB
	})

	r = client
}
