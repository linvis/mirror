package db

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Catalog struct {
	UserID   int `bson:"user_id" json:"user_id"`
	ID       int `bson:"id" json:"id"`
	Label    string
	ParendID int `bson:"parend_id" json:"parend_id"`
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

	cur.Close(context.TODO())

	fmt.Println(res)

	return &res, nil
}
