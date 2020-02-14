package db

import "testing"

func TestInitMongoDB(t *testing.T) {
	InitMongoDB()

	res, _ := GetCatalog(3)
	t.Log(res)

	// NewCatalog(Catalog{
	// 	UserID:   3,
	// 	ID:       1,
	// 	Label:    "level 1",
	// 	ParendID: 0,
	// })
}
