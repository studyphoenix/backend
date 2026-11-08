package model

import "time"

type Flashcard struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	UpdatedAt       time.Time  `json:"updatedAt" example:"2026-12-17T13:47:33Z"`
	UserID          uint       `json:"-"`
	Subtopics       []Subtopic `json:"-" gorm:"many2many:subtopics_flashcards;"`
	Name            string     `json:"name"`
	Content         string     `json:"content"`
	CompletionCount uint       `json:"completionCount"`
}
