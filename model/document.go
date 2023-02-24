package model

type Document struct {
	ID       interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string      `json:"title,omitempty" bson:"title,omitempty"`
	Pages    uint        `json:"pages,omitempty" bson:"pages,omitempty"`
	Language string      `json:"language,omitempty" bson:"language,omitempty"`
	Author   []Author    `json:"author,omitempty" bson:"author,omitempty"`
}

type DocumentWrite struct {
	Title    string   `json:"title,omitempty" bson:"title,omitempty"`
	Pages    uint     `json:"pages,omitempty" bson:"pages,omitempty"`
	Language string   `json:"language,omitempty" bson:"language,omitempty"`
	AuthorID []string `json:"author_id,omitempty" bson:"author_id,omitempty"`
}

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
	Repository interface {
		ListAllAuthor() ([]Author, error)
		GetAuthorDetail(id string) (*Author, error)
		InsertAuthor(author *AuthorWrite) (*Author, error)
		UpdateAuthor(id string, authorWrite AuthorWrite) (*Author, error)
		Delete(id string) error
	}
)
