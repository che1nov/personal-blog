package entity

import "time"

// Article represents the structure of an article
type Article struct {
	ID            string    `json:"id"`             // Unique ID of the article
	Title         string    `json:"title"`          // Title of the article
	Content       string    `json:"content"`        // Content of the article
	PublishedDate time.Time `json:"published_date"` // Date when the article was published
}
