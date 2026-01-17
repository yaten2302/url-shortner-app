package repository

import (
	"context"
	"fmt"
	"time"

	"url-shortner-app/internal/db"
	"url-shortner-app/internal/models"
	"url-shortner-app/pkg/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateShortURL(ctx context.Context, originalURL string) (*models.ShortURL, error) {
	shortCode, err := utils.GenerateShortCode(len(originalURL))
	if err != nil {
		return nil, err
	}
	now := time.Now()

	shortURL := models.ShortURL{
		URL:         originalURL,
		ShortCode:   shortCode,
		CreatedAt:   now,
		UpdatedAt:   now,
		AccessCount: 0,
	}

	if err := Insert(ctx, shortURL); err != nil {
		return nil, err
	}

	return &shortURL, nil
}

// insert shortURL into the database
func Insert(ctx context.Context, shortURL models.ShortURL) error {
	result, err := db.DB.Collection("urls").InsertOne(ctx, shortURL)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with ID: %v\n", result.InsertedID)
	return nil
}

// update shortURL in the database
func Update(ctx context.Context, shortURL models.ShortURL) error {
	_, err := db.DB.Collection("urls").UpdateOne(ctx, map[string]any{"shortCode": shortURL.ShortCode}, map[string]any{
		"$set": map[string]any{
			"url":         shortURL.URL,
			"updatedAt":   shortURL.UpdatedAt,
			"accessCount": shortURL.AccessCount,
		},
	})

	return err
}

// find original url by shortCode
func FindByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	var url models.ShortURL

	err := db.DB.Collection("urls").FindOne(ctx, map[string]any{"shortCode": shortCode}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}

func DeleteByShortCode(ctx context.Context, shortCode string) error {
	_, err := db.DB.Collection("urls").DeleteOne(ctx, map[string]any{"shortCode": shortCode})
	return err
}
