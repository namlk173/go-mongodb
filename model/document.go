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

type DocumentRepository interface {
}
