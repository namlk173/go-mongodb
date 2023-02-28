package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID       interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string      `json:"title,omitempty" bson:"title,omitempty"`
	Pages    uint        `json:"pages,omitempty" bson:"pages,omitempty"`
	Language string      `json:"language,omitempty" bson:"language,omitempty"`
	Author   []Author    `json:"author,omitempty" bson:"author,omitempty"`
}

type DocumentWrite struct {
	Title    string               `json:"title,omitempty" bson:"title,omitempty"`
	Pages    uint                 `json:"pages,omitempty" bson:"pages,omitempty"`
	Language string               `json:"language,omitempty" bson:"language,omitempty"`
	AuthorID []primitive.ObjectID `json:"author_id" bson:"author_id"`
}

type (
	DocumentRepository interface {
		ListAllDocument(skip, limit int64) ([]Document, error)
		InsertDocument(document *DocumentWrite) (interface{}, error)
		GetDocumentDetail(id interface{}) (*Document, error)
		UpdateDocument(objectID interface{}, document *DocumentWrite) error
		DeleteDocument(id interface{}) error
	}
)
