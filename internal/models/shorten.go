package models

import "time"

type ShortURL struct {
	URL         string    `bson:"url" json:"url"`
	ShortCode   string    `bson:"shortCode" json:"shortCode"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt" json:"updatedAt"`
	AccessCount int       `bson:"accessCount" json:"accessCount"`
}
