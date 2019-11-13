package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbConfig string = "root:ironman.abc@/mirrors?charset=utf8&parseTime=True&loc=Local"

type User struct {
	UserID   int    `gorm:"column:user_id"`
	Email    string `gorm:"column:email"`
	Name     string `gorm:"column:user_name"`
	Password string `gorm:"column:passwd"`
	Male     int    `gorm:"column:male"`
	Age      int    `gorm:"column:age"`
	Location string `gorm:"column:location"`
	Role     string `gorm:"column:role"`
	Avatar   string `gorm:"column:avatar"`
}

func (u User) TableName() string {
	return "user"
}

func init() {

	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&User{})

	var count int

	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Email:    "admin@test.com",
			Name:     "admin",
			Password: "admin123",
			Role:     "admin",
			Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		})
	}

	defer db.Close()
}

func QueryUserByName(name string) User {
	db, _ := gorm.Open("mysql", dbConfig)

	var user User
	db.Where("user_name = ?", name).First(&user)

	defer db.Close()

	return user
}
