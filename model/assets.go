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

type CreateAssetRequest struct {
	Name      string `json:"name" example:"Analysis Script" maxLength:"100" minLength:"3" validate:"required" binding:"required,min=3,max=100,printascii"`
	MediaType string `json:"mediaType" example:"application/pdf" minLength:"3" maxLength:"40" validate:"required" binding:"required,contains=/,min=3,max=40,printascii"` // The mediatype (formerly MIME type) of the asset. Will block replacement with asset of another type
	Icon      string `json:"icon" example:"fa-gear" maxLength:"250" minLength:"0" validate:"required" binding:"required,min=3,max=250,printascii"`                       // Must be a valid identifier for vuetify
}

type UpdateAssetRequest struct {
	Name *string `json:"name" example:"Analysis Script" maxLength:"250" minLength:"3" validate:"optional" binding:"omitempty,min=3,max=250,printascii"`
	Icon *string `json:"icon" example:"fa-gear" maxLength:"250" minLength:"0" validate:"optional" binding:"omitempty,min=3,max=250,printascii"` // Must be a valid identifier for vuetify
}
