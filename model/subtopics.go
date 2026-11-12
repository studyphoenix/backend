package model

import "time"

type Subtopic struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"-"`
	UpdatedAt  time.Time   `json:"updatedAt" example:"2026-12-17T13:47:33Z"`
	Name       string      `json:"name" example:"ordinary differential equations"`
	Icon       string      `json:"icon" example:"fa-gear"`
	Progress   uint        `json:"progress" example:"0" maximum:"100" minimum:"0"` // value in [0, 100]
	Assets     []Asset     `json:"assets,omitempty" gorm:"many2many:subtopics_assets;" validate:"optional"`
	Chats      []Chat      `json:"chats,omitempty" gorm:"many2many:subtopics_chats;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" validate:"optional"`
	Subtopics  []Subtopic  `json:"subtopics,omitempty" gorm:"many2many:subtopics_subtopics;" validate:"optional"`
	Flashcards []Flashcard `json:"-" gorm:"many2many:subtopics_flashcards;"`
	SubjectID  uint        `json:"-"`
}

type CreateSubtopicRequest struct {
	Name     string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"required,printascii,min=3,max=100"`
	Icon     string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"required,printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
	Subject  uint   `example:"1" json:"subject" validate:"required"`                                                           // Subject this subtopic belongs to
	Progress uint   `example:"0" json:"progress" validate:"required,min=0,max=100" default:"0" minimum:"0" maximum:"100"`
}

type UpdateSubtopicRequest struct {
	Name     *string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"optional,printascii,min=3,max=100"`
	Icon     *string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"optional,printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
	Subject  *uint   `example:"1" json:"subject" validate:"optional"`                                                           // Subject this subtopic belongs to
	Progress *uint   `example:"0" json:"progress" validate:"optional,min=0,max=100" default:"0" minimum:"0" maximum:"100"`
}
