package model

type Author struct {
	ID        interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string      `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string      `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Address   string      `json:"address,omitempty" bson:"address,omitempty"`
}

type AuthorWrite struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Address   string `json:"address,omitempty" bson:"address,omitempty"`
}

type (
	AuthorRepository interface {
		ListAllAuthor() ([]Author, error)
		GetAuthorDetail(objectID interface{}) (*Author, error)
		InsertAuthor(author *AuthorWrite) (*Author, error)
		UpdateAuthor(author *Author) error
		DeleteAuthor(objectID interface{}) error
	}
)
