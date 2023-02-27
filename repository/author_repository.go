package repository

import (
	"context"
	"go-mongodb/database"
	"go-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type authorRepository struct {
	db *database.Database
}

func NewAuthorRepository(db *database.Database) model.AuthorRepository {
	return &authorRepository{db: db}
}

func (r *authorRepository) ListAllAuthor() ([]model.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var authors []model.Author
	collection := r.db.Client.Database("go-mongodb").Collection("author")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []model.Author{}, err
	}
	defer cur.Close(ctx)

	//for cur.Next(ctx) {
	//	var author model.Author
	//	err = cur.Decode(&author)
	//	if err != nil {
	//		return []model.Author{}, err
	//	}
	//	authors = append(authors, author)
	//}

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
	filter := bson.M{"_id": objectID}
	err := collection.FindOne(ctx, filter).Decode(&author)
	if err != nil {
		return &model.Author{}, err
	}

	return &author, nil
}

func (r *authorRepository) InsertAuthor(author *model.AuthorWrite) (*model.Author, error) {
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
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
