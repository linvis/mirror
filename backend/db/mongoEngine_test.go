package db

import "testing"

func TestInitMongoDB(t *testing.T) {
	InitMongoDB()

	NewCatalog(Catalog{
		UserID:   3,
		ID:       1,
		Label:    "level 1",
		ParendID: 0,
	})

	GetCatalog(3)
}
