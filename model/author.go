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
		GetAuthorDetail(id string) (*Author, error)
		InsertAuthor(author *AuthorWrite) (*Author, error)
		UpdateAuthor(id string, authorWrite AuthorWrite) (*Author, error)
		DeleteAuthor(id string) error
	}
)
