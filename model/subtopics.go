package model

import "time"

type Subtopic struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"-"`
	UpdatedAt  time.Time   `json:"updatedAt" example:"2026-12-17T13:47:33Z"`
	Name       string      `json:"name" example:"ordinary differential equations"`
	Assets     []Asset     `json:"assets,omitempty" gorm:"many2many:subtopics_assets;" validate:"optional"`
	Flashcards []Flashcard `json:"flashcards,omitempty" gorm:"many2many:subtopics_flashcards;" validate:"optional"`
	Chats      []Chat      `json:"chats,omitempty" gorm:"many2many:subtopics_chats;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" validate:"optional"`
	Subtopics  []Subtopic  `json:"subtopics,omitempty" gorm:"many2many:subtopics_subtopics;" validate:"optional"`
	Subjects   []Subject   `json:"-" gorm:"many2many:subtopics_subjects"`
}
