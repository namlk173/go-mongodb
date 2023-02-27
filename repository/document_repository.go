package repository

import (
	"context"
	"go-mongodb/database"
	"go-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type documentRepository struct {
	db *database.Database
}

func NewDocumentRepository(db *database.Database) model.DocumentRepository {
	return &documentRepository{db: db}
}

func (d *documentRepository) ListAllDocument() ([]model.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var documents []model.Document
	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []model.Document{}, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var document model.Document
		err := cur.Decode(&document)
		if err != nil {
			return []model.Document{}, err
		}
		documents = append(documents, document)
	}

	if cur.Err() != nil {
		return []model.Document{}, err
	}

	return documents, nil
}

func (d *documentRepository) GetDocumentDetail(objectID interface{}) (*model.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	var document model.Document
	filter := bson.M{"_id": objectID}
	err := collection.FindOne(ctx, filter).Decode(&document)
	if err != nil {
		return &model.Document{}, err
	}

	return &document, nil
}

func (d *documentRepository) InsertDocument(document *model.DocumentWrite) (*model.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		return &model.Document{}, err
	}

	return &model.Document{
		ID:       res.InsertedID,
		Title:    document.Title,
		Pages:    document.Pages,
		Language: document.Language,
		Author:   document.Author,
	}, nil
}

func (d *documentRepository) UpdateDocument(document *model.Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	filter := bson.M{"_id": document.ID}
	_, err := collection.ReplaceOne(ctx, filter, &document)
	if err != nil {
		return err
	}

	return nil
}

func (d *documentRepository) DeleteDocument(objectID interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	filter := bson.M{"_id": objectID}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
