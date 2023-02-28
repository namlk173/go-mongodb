package repository

import (
	"context"
	"go-mongodb/database"
	"go-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type authorRepository struct {
	db *database.Database
}

func NewAuthorRepository(db *database.Database) model.AuthorRepository {
	return &authorRepository{db: db}
}

func (r *authorRepository) ListAllAuthor(skip, limit int64) ([]model.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var authors []model.Author
	collection := r.db.Client.Database("go-mongodb").Collection("author")
	opt := options.Find().SetLimit(limit).SetSkip(skip)
	cur, err := collection.Find(ctx, bson.D{{"is_deleted", false}}, opt)
	if err != nil {
		return []model.Author{}, err
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &authors); err != nil {
		return []model.Author{}, err
	}

	if err = cur.Err(); err != nil {
		return []model.Author{}, err
	}

	return authors, nil
}

func (r *authorRepository) GetAuthorDetail(objectID interface{}) (*model.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := r.db.Client.Database("go-mongodb").Collection("author")
	var author model.Author
	filter := bson.M{"_id": objectID, "is_deleted": false}
	err := collection.FindOne(ctx, filter).Decode(&author)
	if err != nil {
		return &model.Author{}, err
	}

	return &author, nil
}

func (r *authorRepository) InsertAuthor(author *model.AuthorWrite) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := r.db.Client.Database("go-mongodb").Collection("author")
	res, err := collection.InsertOne(ctx, author)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID, nil
}

func (r *authorRepository) UpdateAuthor(author *model.Author) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := r.db.Client.Database("go-mongodb").Collection("author")
	filter := bson.M{"_id": author.ID}
	_, err := collection.ReplaceOne(ctx, filter, author)
	if err != nil {
		return err
	}

	return nil
}

func (r *authorRepository) DeleteAuthor(objectID interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	collection := r.db.Client.Database("go-mongodb").Collection("author")
	filter := bson.M{"_id": objectID}
	updateQuery := bson.D{{"$set", bson.D{{"is_deleted", true}}}}
	_, err := collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return err
	}

	return nil
}
