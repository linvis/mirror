package db

import "errors"

type Document struct {
	DocID  string `xorm:"doc_id" json:"doc_id"`
	UserID int    `xorm:"user_id" json:"-"`
	Raw    string `xorm:"MEDIUMTEXT"`
	HTML   string `xorm:"MEDIUMTEXT html" json:"html"`
}

// TableName return table name
func (d Document) TableName() string {
	return "document"
}

func NewDocument(doc *Document) error {
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

	_, err := x.Where("user_id = ? AND doc_id = ?", user_id, doc_id).Get(&docs)
	if err != nil {
		return nil, errors.New("no document")
	}

	return &docs, nil
}
