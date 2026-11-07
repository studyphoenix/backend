package model

import "time"

type Asset struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"-"`
	Name      string     `json:"name" example:"my asset name"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2027-01-01T00:00:00Z"`
	MediaType string     `json:"mediaType" example:"application/pdf"`
	Icon      string     `json:"icon" example:"fa-gear"`
	Subtopics []Subtopic `json:"-" gorm:"many2many:subtopics_assets;"`
	Subjects  []Subject  `json:"-" gorm:"many2many:subjects_assets;"`
}
