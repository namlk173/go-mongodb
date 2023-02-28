package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	IsDeleted bool               `json:"is_deleted" bson:"is_deleted"`
}

type AuthorWrite struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Address   string `json:"address,omitempty" bson:"address,omitempty"`
	IsDeleted bool   `json:"is_deleted" bson:"is_deleted"`
}

type (
	AuthorRepository interface {
		ListAllAuthor(skip, limit int64) ([]Author, error)
		GetAuthorDetail(objectID interface{}) (*Author, error)
		InsertAuthor(author *AuthorWrite) (interface{}, error)
		UpdateAuthor(author *Author) error
		DeleteAuthor(objectID interface{}) error
	}
)
