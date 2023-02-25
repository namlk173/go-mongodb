package controller

import (
	"context"
	"go-mongodb/database"
	"go-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) model.Repository {
	return &repository{db: db}
}

func (r *repository) ListAllAuthor() ([]model.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var authors []model.Author
	collection := r.db.Client.Database("go-mongodb").Collection("author")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []model.Author{}, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var author model.Author
		err = cur.Decode(&author)
		if err != nil {
			return []model.Author{}, err
		}
		authors = append(authors, author)
	}
	if err = cur.Err(); err != nil {
		return []model.Author{}, err
	}

	return authors, nil
}

func (r *repository) GetAuthorDetail(id string) (*model.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &model.Author{}, err
	}

	collection := r.db.Client.Database("go-mongodb").Collection("author")
	var author model.Author
	filter := bson.M{"_id": objectID}
	err = collection.FindOne(ctx, filter).Decode(&author)
	if err != nil {
		return &model.Author{}, nil
	}

	return &author, nil
}

func (r *repository) InsertAuthor(author *model.AuthorWrite) (*model.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	collection := r.db.Client.Database("go-mongodb").Collection("author")
	res, err := collection.InsertOne(ctx, author)
	if err != nil {
		return &model.Author{}, nil
	}

	return &model.Author{
		ID:        res.InsertedID,
		FirstName: author.FirstName,
		LastName:  author.LastName,
		Address:   author.Address,
	}, nil
}

func (r *repository) UpdateAuthor(id string, authorWrite model.AuthorWrite) (*model.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &model.Author{}, err
	}

	collection := r.db.Client.Database("go-mongodb").Collection("author")
	updateQuery := bson.D{{"$set", bson.D{
		{"first_name", authorWrite.FirstName},
		{"last_name", authorWrite.LastName},
		{"address", authorWrite.Address},
	}}}
	_, err = collection.UpdateByID(ctx, objectID, updateQuery)
	if err != nil {
		return &model.Author{}, nil
	}

	return r.GetAuthorDetail(id)
}

func (r *repository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := r.db.Client.Database("go-mongodb").Collection("author")
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
