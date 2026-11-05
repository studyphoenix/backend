package model

import "time"

type Asset struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UpdatedAt time.Time `json:"updatedAt" example:"2027-01-01T00:00:00Z"`
	MediaType string    `json:"mediaType" example:"application/pdf"`
	// TODO: Many to many relationship with subtopics and subjects
}
