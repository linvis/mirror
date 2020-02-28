package db

import (
	"errors"
)

type Document struct {
	Id     int64  `json:"id"`
	DocKey string `xorm:"doc_key" json:"doc_key"`
	UserID int    `xorm:"user_id" json:"-"`
	Raw    string `xorm:"MEDIUMTEXT"`
	HTML   string `xorm:"MEDIUMTEXT html" json:"html"`
}

// TableName return table name
func (d Document) TableName() string {
	return "document"
}

func NewDocument(doc *Document) error {
	oldRec := Document{}
	if doc.Id != 0 {
		has, err := x.Where("id = ?", doc.Id).Get(&oldRec)
		if err != nil {
			return err
		} else if has {
			x.Update(doc, &Document{Id: oldRec.Id})
			return nil
		}
	}

	_, err := x.Insert(doc)

	return err
}

func GetAllDocument(user_id int) (*[]Document, error) {
	docs := &[]Document{}

	err := x.Where("user_id = ?", user_id).Find(docs)
	if err != nil {
		return nil, errors.New("no document")
	}

	return docs, nil
}

func GetDocumentByID(user_id int, doc_id string) (*Document, error) {
	docs := Document{}

	_, err := x.Where("user_id = ? AND id = ?", user_id, doc_id).Get(&docs)
	if err != nil {
		return nil, errors.New("no document")
	}

	return &docs, nil
}
