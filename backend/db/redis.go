package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

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

func SetToRedis(key string, val interface{}, timeout time.Duration) error {
	return r.Set(key, val, timeout).Err()
}

func GetFromRedis(key string) (string, error) {
	return r.Get(key).Result()
}

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
