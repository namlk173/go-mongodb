package repository

import (
	"go-mongodb/database"
	"go-mongodb/model"
)

type documentRepository struct {
	db *database.Database
}

func NewDocumentRepository(db *database.Database) model.DocumentRepository {
	return &documentRepository{db: db}
}
