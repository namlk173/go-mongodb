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
	Author   []Author `json:"author,omitempty" bson:"author,omitempty"`
}

type (
	DocumentRepository interface {
		ListAllDocument() ([]Document, error)
		InsertDocument(document *DocumentWrite) (*Document, error)
		GetDocumentDetail(id interface{}) (*Document, error)
		UpdateDocument(document *Document) error
		DeleteDocument(id interface{}) error
	}
)
