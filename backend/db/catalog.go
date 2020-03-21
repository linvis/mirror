package db

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	FileTypeNoteBook = "nb"
	FileTypeMarkdown = "md"
)

type Reminder struct {
	Enable   bool  `json:"enable"`
	Count    int   `json:"count"`
	LastTime int64 `bson:"last_time" json:"last_time"`
	NextTime int64 `bson:"next_time" json:"next_time"`
}

type MetaData struct {
	Title    string   `json:"title"`
	Level    int      `json:"level"`
	FileType string   `bson:"filetype" json:"filetype"`
	Tag      []string `json:"tag"`
}

type Catalog struct {
	UserID   int       `bson:"user_id" json:"-"`
	ID       int       `bson:"id" json:"id"`
	Key      string    `bson:"key" json:"key"`
	Parent   string    `bson:"parent" json:"parent"`
	Meta     MetaData  `bson:"metadata" json:"metadata"`
	Remind   Reminder  `bson:"reminder" json:"reminder"`
	Children []Catalog `bson:"children" json:"children"`
}

func UpdateCatalog(catalog Catalog) error {
	var res Catalog
	filter := bson.M{
		"user_id": catalog.UserID,
	}

	update := bson.D{
		{"$set", catalog},
	}

	err := collections.FindOne(context.TODO(), filter).Decode(&res)
	if err == mongo.ErrNoDocuments {
		_, err = collections.InsertOne(context.TODO(), catalog)
		if err != nil {
			log.Warn(err)
			return err
		}
	} else if err != nil {
		log.Warn(err)
		return err
	}

	_, err = collections.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func NewDefaultCatalog(user_id int) (Catalog, error) {
	catalog := Catalog{
		UserID: user_id,
		ID:     0,
		Key:    "0",
		Meta: MetaData{
			Title:    "Home",
			Level:    0,
			FileType: FileTypeNoteBook,
		},
		Remind: Reminder{
			Enable:   false,
			Count:    0,
			LastTime: 0,
			NextTime: 0,
		},
		Parent:   "0",
		Children: []Catalog{},
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
