package db

import "errors"

type User struct {
	UserID   int    `xorm:"not null 'user_id'"`
	Email    string `xorm:"char(32) not null 'email'"`
	Name     string `xorm:"char(32) not null 'user_name'"`
	Password string `xorm:"char(32) not null 'passwd'"`
	Male     int    `xorm:"tinyint(4) 'male'"`
	Age      int    `xorm:"tinyint(4) 'age'"`
	Location string `xorm:"char(32) 'location'"`
	Role     string `xorm:"char(32) 'role'"`
	Avatar   string `xorm:"char(255) 'avatar'"`
}

func (u User) TableName() string {
	return "user"
}

// func init() {
// 	db, err := database.GetDB()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	db.AutoMigrate(&User{})

// 	var count int

// 	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
// 		db.Create(&User{
// 			Email:    "admin@test.com",
// 			Name:     "admin",
// 			Password: "admin123",
// 			Role:     "admin",
// 			Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
// 		})
// 	}
// }

func GetUserByName(name string) (*User, error) {
	// db, err := database.GetDB()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// var user User

	// db.Where("user_name = ?", name).First(&user)

	u := new(User)

	has, err := x.Where("user_name = ?", name).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("no user")
	}

	return u, nil
}

func GetUserByID(id int) (*User, error) {

	u := new(User)

	has, err := x.Where("user_id = ?", id).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("no user")
	}

	return u, nil
}
