package model

import "time"

type Subject struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"-"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2026-12-17T13:47:33Z"`
	Name      string     `json:"name" example:"Analysis"`
	Icon      string     `json:"icon" example:"fa-gear"`
	Assets    []Asset    `json:"assets,omitempty" gorm:"many2many:subjects_assets;" validate:"optional"`
	Subtopics []Subtopic `json:"-" gorm:"many2many:subtopics_subjects"`
}
