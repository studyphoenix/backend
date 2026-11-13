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

type CreateFlashcardRequest struct {
	Name    string `example:"Analysis" json:"name" maxLength:"100" minLength:"3" validate:"required" binding:"required,min=3,max=100,printascii"`
	Content string `example:"summarized content for this flashcard with everything needed for the model to ask sensible questions" json:"content" validate:"required" binding:"required,min=0"`
}

type UpdateFlashcardRequest struct {
	Name    *string `example:"ODEs" json:"name" maxLength:"100" minLength:"3" validate:"optional" binding:"omitempty,min=3,max=100"`
	Content *string `example:"summarized content for this flashcard with everything needed for the model to ask sensible questions" json:"content" validate:"optional" binding:"omitempty,min=0"`
}
