package model

import "time"

type Subtopic struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"-"`
	UpdatedAt  time.Time   `json:"updatedAt" example:"2026-12-17T13:47:33Z"`
	Name       string      `json:"name" example:"ordinary differential equations"`
	Icon       string      `json:"icon" example:"fa-gear"`
	Assets     []Asset     `json:"assets,omitempty" gorm:"many2many:subtopics_assets;" validate:"optional"`
	Chats      []Chat      `json:"chats,omitempty" gorm:"many2many:subtopics_chats;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" validate:"optional"`
	Subtopics  []Subtopic  `json:"subtopics,omitempty" gorm:"many2many:subtopics_subtopics;" validate:"optional"`
	Flashcards []Flashcard `json:"-" gorm:"many2many:subtopics_flashcards;"`
	Subjects   []Subject   `json:"-" gorm:"many2many:subtopics_subjects"`
}

type CreateSubtopicRequest struct {
	Name string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"required,printascii,min=3,max=100"`
	Icon string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"required,printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
}

type UpdateSubtopicRequest struct {
	Name *string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"optional,printascii,min=3,max=100"`
	Icon *string `example:"fa-gear" json:"icon" maxLength:"100" minLength:"0" validate:"optional,printascii,min=0,max=100"` // The Icon of the subject. Must be an icon that can be rendered in the browser
}
