package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var (
	x      *xorm.Engine
	tables []interface{}
)

func loadDBConfig(file string) (*DBConfig, error) {
	var config DBConfig

	dbFile, err := os.Open(file)
	defer dbFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	jsonParser := json.NewDecoder(dbFile)
	jsonParser.Decode(&config)

	return &config, nil
}

func SetEngine(config *DBConfig) (*xorm.Engine, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Host, config.Port, config.Database)

	engine, err := xorm.NewEngine("mysql", connStr)

	return engine, err
}

func GetEngine() (*xorm.Engine, error) {
	if x == nil {
		return nil, errors.New("invalid engine")
	}

	return x, nil
}

func InitDB(path string) {

	tables = append(tables, new(User), new(DailyEvt), new(DailyEvtType))

	config, err := loadDBConfig(path)
	// config, err := loadDBConfig("db.config")
	if err != nil {
		return
	}

	engine, err := SetEngine(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = engine.Sync2(tables...)
	if err != nil {
		fmt.Println(err)
		return
	}

	x = engine
}
