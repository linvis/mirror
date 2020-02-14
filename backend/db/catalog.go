package db

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Catalog struct {
	UserID   int    `bson:"user_id" json:"-"`
	ID       int    `bson:"id" json:"id"`
	Label    string `json:"label"`
	ParendID int    `bson:"parend_id" json:"parend_id"`
}

func NewCatalog(catalog Catalog) error {
	var res Catalog
	filter := bson.M{
		"user_id": catalog.UserID,
		"id":      catalog.ID,
	}

	err := collections.FindOne(context.TODO(), filter).Decode(&res)
	if err == nil {
		log.Fatal("duplicated catalog")
		return errors.New("duplicated catalog")
	} else if err != mongo.ErrNoDocuments {
		log.Fatal(err)
		return err
	}

	_, err = collections.InsertOne(context.TODO(), catalog)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func NewDefaultCatalog(user_id int) (Catalog, error) {
	catalog := Catalog{
		UserID:   user_id,
		ID:       1,
		Label:    "默认笔记本",
		ParendID: 0,
	}

	_, err := collections.InsertOne(context.TODO(), catalog)

	return catalog, err
}

func GetCatalog(user_id int) (*[]Catalog, error) {
	filter := bson.M{
		"user_id": user_id,
	}

	findOptions := options.Find()

	var res []Catalog

	cur, err := collections.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem Catalog

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.TODO())

	if len(res) == 0 {
		// create default catalog
		catalog, err := NewDefaultCatalog(user_id)

		res = append(res, catalog)

		return &res, err
	}

	return &res, nil
}
