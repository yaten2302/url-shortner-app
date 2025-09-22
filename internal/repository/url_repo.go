package repository

import (
	"context"
	"fmt"

	"url-shortner-app/internal/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type URLRepository struct {
	collection *mongo.Collection
}

func NewURLRepository(col *mongo.Collection) *URLRepository {
	return &URLRepository{collection: col}
}

// insert shortURL into the database
func (r *URLRepository) Insert(ctx context.Context, shortURL models.ShortURL) error {
	result, err := r.collection.InsertOne(ctx, shortURL)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with ID: %v\n", result.InsertedID)
	return nil
}

// update shortURL in the database
func (r *URLRepository) Update(ctx context.Context, shortURL models.ShortURL) error {
	_, err := r.collection.UpdateOne(ctx, map[string]any{"shortCode": shortURL.ShortCode}, map[string]any{
		"$set": map[string]any{
			"url":          shortURL.URL,
			"updatedAt":   shortURL.UpdatedAt,
			"accessCount": shortURL.AccessCount,
		},
	})

	return err
}

// find original url by shortCode
func (r *URLRepository) FindByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	var url models.ShortURL

	err := r.collection.FindOne(ctx, map[string]any{"shortCode": shortCode}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}
