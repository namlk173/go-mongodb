package repository

import (
	"context"
	"errors"
	"go-mongodb/database"
	"go-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type documentRepository struct {
	db *database.Database
}

func NewDocumentRepository(db *database.Database) model.DocumentRepository {
	return &documentRepository{db: db}
}

func (d *documentRepository) ListAllDocument(skip, limit int64) ([]model.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var documents []model.Document
	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	aggregationOption := bson.D{
		{
			"$lookup", bson.D{
				{"from", "author"},
				{"localField", "author_id"},
				{"foreignField", "_id"},
				{"as", "author"},
			},
		},
	}
	unsetFieldOption := bson.D{
		{
			"$project", bson.D{
				{"author_id", 0},
			},
		},
	}

	skipOption := bson.D{
		{
			"$skip", skip,
		},
	}

	limitOption := bson.D{
		{
			"$limit", limit,
		},
	}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{aggregationOption, unsetFieldOption, skipOption, limitOption})
	if err != nil {
		return []model.Document{}, err
	}
	defer cur.Close(ctx)

	//Create a new array contain authors who's is_deleted = false - not be deleted.
	var existAuthor []model.Author
	for cur.Next(ctx) {
		var document model.Document
		err := cur.Decode(&document)
		if err != nil {
			return []model.Document{}, err
		}

		// Filter authors have is_deleted = false and remove authors have is_deleted field = true
		// Using golang to modifier Author in document
		// Need a sub query option to get author didn't be deleted.
		for _, author := range document.Author {
			if !author.IsDeleted {
				existAuthor = append(existAuthor, author)
			}
		}

		document.Author = existAuthor
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
	matchIdOption :=
		bson.D{
			{"$match", bson.D{
				{"_id", objectID},
			}},
		}
	aggregationOption := bson.D{
		{
			"$lookup", bson.D{
				{"from", "author"},
				{"localField", "author_id"},
				{"foreignField", "_id"},
				{"as", "author"},
			},
		},
	}
	unsetFieldOption := bson.D{
		{
			"$project", bson.D{
				{"author_id", 0},
			},
		},
	}

	cur, err := collection.Aggregate(ctx, mongo.Pipeline{matchIdOption, aggregationOption, unsetFieldOption})
	defer cur.Close(ctx)

	if err != nil {
		return &model.Document{}, err
	}

	for cur.Next(ctx) {
		err = cur.Decode(&document)
		if err != nil {
			return &model.Document{}, err
		}
	}

	if document.ID == nil {
		return &model.Document{}, errors.New("document not found")
	}

	var existAuthor []model.Author
	for _, author := range document.Author {
		if !author.IsDeleted {
			existAuthor = append(existAuthor, author)
		}
	}
	document.Author = existAuthor
	return &document, nil
}

// InsertDocument Return objectID of document
func (d *documentRepository) InsertDocument(document *model.DocumentWrite) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (d *documentRepository) UpdateDocument(objectID interface{}, document *model.DocumentWrite) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := d.db.Client.Database("go-mongodb").Collection("documents")
	filter := bson.M{"_id": objectID}
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
