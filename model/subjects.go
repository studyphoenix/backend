package model

import "time"

type Subject struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"-"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2026-12-17T13:47:33Z"`
	Name      string     `json:"name" example:"Analysis"`
	Icon      string     `json:"icon" example:"fa-gear"`
	Assets    []Asset    `json:"assets,omitempty" gorm:"many2many:subjects_assets;" validate:"optional"`
	Subtopics []Subtopic `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}

type CreateSubjectRequest struct {
	Name string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"required" binding:"required,printascii,min=3,max=100"`
	Icon string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"required" binding:"required,printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
}

type UpdateSubjectRequest struct {
	Name *string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"optional" binding:"printascii,min=3,max=100"`
	Icon *string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"optional" binding:"printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
}
